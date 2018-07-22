package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
