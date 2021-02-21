package main

import "net/http"

func main() {
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":3000", nil)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloworld!!"))
}
