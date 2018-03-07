package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Host == "muffins.cupcakes.org" {
		log.Print("Sweet, sweet muffins")
		fmt.Fprint(w, "Hello muffin lover")
	} else {
		w.WriteHeader(404)
		log.Print("Correct header not found, no muffins for you :(")
		fmt.Fprint(w, "Wot no muffins :(")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
