package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jeonjonghyeok/todolist_ex2/db"
)

func must(err error) {
	if err == db.ErrorNotFound {
		log.Println("DB Not Found")
		panic(notFoundError)
	} else if err != nil {
		log.Println("Internal Error")
		panic(internalError)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(w).Encode(v))

}

func parseJSON(r io.Reader, v interface{}) interface{} {
	json.NewDecoder(r).Decode(v)
	return v
}
