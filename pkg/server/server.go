package server

import (
	"fmt"
	"net/http"
	"os"

	"gop/pkg/api"
)

func Run() error {
	api.Init()
	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("TODO_PORT")), nil)
}
