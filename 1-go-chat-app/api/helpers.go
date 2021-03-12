package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func parseJSON(r io.Reader, v interface{}) {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		log.Println("parse json", err)
		must(err)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(w).Encode(v))
}
