package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir("tmp"))))
	// http.Handle("/", http.FileServer(http.Dir("tmp")))
	// http.Handle("/tmp", http.FileServer(http.Dir("tmp")))
	http.HandleFunc("/mypic", getmypic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getmypic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/c.jpg")
}
