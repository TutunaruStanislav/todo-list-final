package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"gop/pkg/api"

	"github.com/go-chi/chi/v5"
)

func Run(db *sql.DB) error {
	router := chi.NewRouter()
	api.Init(db, router)
	fileServer := http.FileServer(http.Dir("web"))
	router.Handle("/*", fileServer)

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("TODO_PORT")), router)
}
