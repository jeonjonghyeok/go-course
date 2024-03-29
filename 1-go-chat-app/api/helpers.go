package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/db"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/token"
)

func must(err error) {
	if err == db.ErrUnauthorized {
		panic(unauthorizedError)
	} else if err != nil {
		panic(internalError)
	}
}
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

func userID(r *http.Request) int {
	t := r.URL.Query().Get("token")
	id, err := token.Parse(t)
	if err != nil {
		log.Println(err)
		panic(unauthorizedError)
	}
	return id
}

func parseIntParam(r *http.Request, key string) int {
	vars := mux.Vars(r)
	if v, ok := vars[key]; ok {
		id, err := strconv.Atoi(v)
		if err == nil {
			return id
		}
	}
	panic(malformedInputError)
}
