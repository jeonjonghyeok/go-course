package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jeonjonghyeok/todolist_ex2/db"
)

func must(err error) {
	if err == db.ErrorNotFound {
		log.Println("DB Not Found")
		panic(notFoundError)
	} else if err != nil {
		log.Println("Internal Error", err)
		panic(internalError)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(w).Encode(v))

}

func parseJSON(r io.Reader, v interface{}) {
	json.NewDecoder(r).Decode(v)
}

func parseIntParam(r *http.Request, key string) int {
	vars := mux.Vars(r)
	if v, ok := vars[key]; ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	panic(malformedInputError)
}
