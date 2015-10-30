package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", greeting)
	http.ListenAndServe(":3000", nil)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I am ready to be a Web server. "))
}
