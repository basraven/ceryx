package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"ceryx/handlers/qbittorent"
	"ceryx/senders/telegram"
)

func main() {
	port := "80"
	log.Printf("Starting Ceryx on port %s\n", port)

	sender := telegram.Send

	r := mux.NewRouter()
	r.HandleFunc(qbittorent.UrlSuffix, func (w http.ResponseWriter, r *http.Request) {
		qbittorent.Handle(w,r,sender)
	})

	http.ListenAndServe(":80", r)
}