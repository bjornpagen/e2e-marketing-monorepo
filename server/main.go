package main

import (
	"log"
	"net/http"

	"github.com/bjornpagen/e2e-marketing-monorepo/server/hashlookup"
)

func main() {
	hashToEmailMap := map[string]string{
		"b1a6b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4": "bjorn.pagen@example.com",
	}

	hashlookupClient := hashlookup.New(hashToEmailMap)
	http.HandleFunc("/hashlookup", hashlookupClient.HashLookupHandler)

	log.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
