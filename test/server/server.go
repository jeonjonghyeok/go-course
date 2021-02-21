package main

import (
	"net/http"
)

func main() {
	/*
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World"))
		})
		http.ListenAndServe(":5000", nil)
	*/
	http.Handle("/", new(testHandler))
	http.ListenAndServe(":5000", nil)

}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Reqeust Path is " + req.URL.Path
	w.Write([]byte(str))
}
