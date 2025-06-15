package api

import (
	"net/http"
	"os"
)

type SignInResponse struct {
	Token string `json:"token"`
}

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

func signInHandler(w http.ResponseWriter, r *http.Request) {
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
		writeError(w, "wrong password provided", http.StatusBadRequest)
		return
	}

	token, err := createToken(user.Username)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, SignInResponse{Token: token}, http.StatusOK)
}
