package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"gop/pkg/api"
)

func Run(db *sql.DB) error {
	api.Init(db)
	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("TODO_PORT")), nil)
}
