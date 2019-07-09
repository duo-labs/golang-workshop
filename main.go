package main

import "net/http"

// HelloHandler is an example handler which returns the response "Hello"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
