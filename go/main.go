package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func blisteur(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Blisteur!")
}

func id(w http.ResponseWriter, r *http.Request) {
	r.
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/blisteur", blisteur)
	http.HandleFunc("/", id)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
