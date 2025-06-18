package api

import (
	"net/http"
	"os"
)

type SignInResponse struct {
	Token string `json:"token"`
}

// auth - middleware handler to check authorization.
//
// It reads the value of Cookie[“token”] header and if it was passed and is valid - passes control to the next handler, otherwise authentication error.
func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := os.Getenv("TODO_PASSWORD")
		if len(password) > 0 {
			var jwt string
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}

			valid := verifyToken(jwt)
			if !valid {
				writeError(w, "authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

// SignInHandler is the POST request handler for the /api/singnin request.

// It receives, validates and compares the value of the password parameter
// with the preset password in the TODO_PASSWORD environment variable,
// and if the password is correct, it returns the generated jwt-token, otherwise an error.
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	user, err := validatePassword(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	password := os.Getenv("TODO_PASSWORD")
	if len(password) == 0 {
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if user.Password != password {
		writeError(w, "wrong password provided", http.StatusUnauthorized)
		return
	}

	token, err := createToken(user.Username)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, SignInResponse{Token: token}, http.StatusOK)
}
