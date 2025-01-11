package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleError(err Error, w http.ResponseWriter) bool {
	if err.Err != nil {
		log.Println(err.Err)
		res_err := struct {
			Message string
			Code    int
		}{http.StatusText(err.Code), err.Code}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(res_err)
		return true
	}
	return false
}
