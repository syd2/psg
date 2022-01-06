package utils

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error, something went wrong!!": err.Error()})
		return
	}
}
