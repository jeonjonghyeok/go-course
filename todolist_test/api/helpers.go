package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeonjonghyeok/todolist_test/db"
)

func must(err error) {
	if err == db.ErrNotFound {
		log.Println("not found")
		panic(notFoundError)
	} else if err != nil {
		log.Println("internal error: ", err)
		panic(internalError)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
