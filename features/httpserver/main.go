package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello main page "))
}
func main() {
	http.Handle("/", http.HandlerFunc(home))
	log.Println("starting server on : 4000")

	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}
