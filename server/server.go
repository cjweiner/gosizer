package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Image struct {
	Data      []byte
	Name      string
	Extension string
}

func OriginalImageHandler(w http.ResponseWriter, r *http.Request) {

}

func ResizedImageHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	log.Println("Starting Server...")
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../public/")))
	router.HandleFunc("/api/image", OriginalImageHandler).Methods("POST")
	router.HandleFunc("/api/image", ResizedImageHandler).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
