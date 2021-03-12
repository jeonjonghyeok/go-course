package api

import (
	"encoding/json"
	"net/http"
)

func must(e error) {
	panic(e)
}

type apiErrorWriter interface {
	Write(w http.ResponseWriter)
}

func handlePanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if c, ok := r.(apiErrorWriter); ok {
					c.Write(w)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})

}

type simpleError struct {
	message string
	status  int
}

func (e simpleError) Write(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(e.message)
}
