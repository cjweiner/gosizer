package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Server...")
	http.Handle("/", http.FileServer(http.Dir("../public/")))
	log.Println("Server listening on 8888")
	http.ListenAndServe(":8888", nil)
}
