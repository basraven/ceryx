package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	"github.com/basraven/ceryx/internal/handlers/qbittorent"
	"github.com/basraven/ceryx/internal/senders/telegram"
)

func main() {
	port := "80"
	log.Printf("Starting Ceryx on port %s\n", port)

	sender := telegram.Send

	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Please use one of the configured handlers!\n")
	})
	r.HandleFunc(qbittorent.UrlSuffix, func (w http.ResponseWriter, r *http.Request) {
		qbittorent.Handle(w,r,sender)
		fmt.Fprintf(w, "Message send!\n")
	})

	http.ListenAndServe(":80", r)
}