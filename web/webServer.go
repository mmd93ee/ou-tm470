// Package web used to create a Web Server to place Handler logic
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", defaultHandler) // Each call to "/" will invoke defaultHandler
	http.HandleFunc("/blockchain/view/", blockchainViewHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Default handler to catch-all
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Default Handler called from %s.  Please try alternative methods such as /blockchain/", r.RemoteAddr)
}

// Handler to manage requests to /blockchain/ subchain
func blockchainViewHandler(w http.ResponseWriter, r *http.Request) {

	// Take the URL beyond /blockchain/ and split into request and value strings
	requestAction := strings.Split(r.URL.String(), "/")
	requestItem := &requestAction[3]

	// DEBUG
	fmt.Fprintf(w, "Request made for item ID %v", requestItem)

	if requestItem == nil { //Request item is invalid so display that blockID only
		fmt.Printf("Stub behaviour - no block ID hence print entre chain")
	} else {
		fmt.Fprintf(w, "Stub behaviour - print block number %s", requestAction)
	}
}