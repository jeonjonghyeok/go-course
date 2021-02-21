package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/learningspoons-go/todolist/db"
)

func must(err error) {
	if err == db.ErrNotFound {
		log.Println("not found:", err)
		panic(notFoundError)
	} else if err != nil {
		log.Println("internal error:", err)
		panic(internalError)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(w).Encode(v))
}

func parseJSON(r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		log.Println("parsing json body:", err)
		panic(malformedInputError)
	}
}

/*
func parseIntParam(r *http.Request,key string) int {
	vars
}
*/
