package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var congress []Congressman

func handleCongressman(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for i, c := range congress {
		if c.Id == id {
			log.Printf("Found at %v", i)
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	fmt.Fprintf(w, "- not found -")
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	congress = make([]Congressman, 2)
	congress[0] = Congressman{Id: "pr", Name: "Peter Russo"}
	congress[1] = Congressman{Id: "fu", Name: "Frank Underwood"}

	http.HandleFunc("/", handleCongressman)
	http.ListenAndServe(":8090", nil)
}
