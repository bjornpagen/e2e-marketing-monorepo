package main

import (
	"log"
	"net/http"

	"github.com/bjornpagen/e2e-marketing-monorepo/server/idlookup"
)

func main() {
	idToEmailMap := map[string]string{
		"b1a6b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4": "bjorn.pagen@example.com",
	}

	idlookupClient := idlookup.New(idToEmailMap)
	http.HandleFunc("/idlookup", idlookupClient.IdLookupHandler)

	log.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
