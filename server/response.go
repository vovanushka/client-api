package server

import (
	"net/http"
)

func Error(w http.ResponseWriter, code int, message string) {
	Json(w, code, []byte("{\"error\":\""+message+"\"}"))
}

func Json(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
