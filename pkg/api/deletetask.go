package api

import (
	"net/http"

	"gop/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DeleteTask(id)
	if err != nil {
		writeError(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJson(w, SuccessResponse, http.StatusOK)
}
