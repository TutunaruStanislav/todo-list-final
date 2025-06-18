package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"gop/pkg/api"

	"github.com/go-chi/chi/v5"
)

// Run creates a new chi router, initializes the api handlers as well as
// the static file handler from the web/ directory, and starts the server.

// If successful, the server will start on the port you specify
// in the TODO_PORT environment variable, otherwise it will return an error.
func Run(db *sql.DB) error {
	router := chi.NewRouter()
	api.Init(db, router)
	fileServer := http.FileServer(http.Dir("web"))
	router.Handle("/*", fileServer)

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("TODO_PORT")), router)
}
