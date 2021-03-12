package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ChatAPI() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/rooms", getRooms).Methods(http.MethodGet, http.MethodOptions)

	return router
}
