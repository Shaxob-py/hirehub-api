package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func ResponseError(w http.ResponseWriter, statusCode int, err string) {
	ResponseWithJson(w, statusCode, map[string]string{"error": err})
}
