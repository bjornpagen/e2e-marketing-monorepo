package main

import (
	"log"
	"net/http"

	"github.com/bjornpagen/e2e-marketing-monorepo/server/hashlookup"
)

func main() {
	http.HandleFunc("/hash", hashlookup.HashLookupHandler)

	log.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
