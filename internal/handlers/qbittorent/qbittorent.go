package qbittorent


import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Request path:	http://ceryx.bas/qbittorent/%N/%L/%D/%C/%Z"
var UrlSuffix = "/qbittorent/{name}/{category}/{filecount}/{size}"

func Handle(w http.ResponseWriter, r *http.Request, cb func(string) ) {
	getVars := mux.Vars(r)
	torName := getVars["name"]
	torCategory := getVars["category"]
	torfileCount := getVars["filecount"]
	torSize := getVars["size"]


	log.Printf("New message on Torrent: %s\n", torName)
	
	cb(composeMessage(torName, torCategory, torfileCount, torSize))
}

func composeMessage(torName string, torCategory string, torfileCount string, torSize string) (string) {
	return "QBIT: Completed " + torName + "\n" + torCategory + " " + torfileCount + " files with size " + torSize
}