package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/bjornpagen/e2e-marketing-monorepo/server/idlookup"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using environment variables instead.")
	}

	// check for required environment variables
	if os.Getenv("ID_DB") == "" {
		log.Fatalln("ID_DB environment variable not set.")
	}
}

func main() {
	// idToEmailMap := map[string]string{
	// 	"b1a6b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4b0c4": "bjorn.pagen@example.com",
	// }

	// load the map from a file
	loadEnv()

	// open the file
	file, err := os.Open(os.Getenv("ID_DB"))
	if err != nil {
		log.Fatalln("Error opening ID_DB file:", err)
	}

	// read the whole file's contents into a byte slice
	var idDb []byte
	idDb, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading ID_DB file:", err)
	}

	var idToEmailMap map[string]string
	err = json.Unmarshal(idDb, &idToEmailMap)
	if err != nil {
		log.Fatalln("Error unmarshalling ID_DB file:", err)
	}

	// sample file:
	//

	// TODO: ip rate limiting
	idlookupClient := idlookup.New(idToEmailMap)
	http.HandleFunc("/idlookup", idlookupClient.IdLookupHandler)

	log.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
