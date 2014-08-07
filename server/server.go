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

func GetResizedImage(w http.ResponseWriter, r *http.Request) {

}

func main() {
	log.Println("Starting Server...")
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../public/")))
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
