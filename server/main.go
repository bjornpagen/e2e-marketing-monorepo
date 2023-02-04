package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/joho/godotenv"

	"github.com/bjornpagen/e2e-marketing-monorepo/server/lookup"
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

func loadDb() (map[lookup.Id]lookup.User, error) {
	db := make(map[lookup.Id]lookup.User)

	b, err := ioutil.ReadFile(os.Getenv("ID_DB"))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &db); err != nil {
		return nil, err
	}

	return db, nil
}

func setupRouter() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(middleware.Logger)

	// Enable httprate request limiter of 100 requests per minute.
	//
	// In the code example below, rate-limiting is bound to the request IP address
	// via the LimitByIP middleware handler.
	//
	// To have a single rate-limiter for all requests, use httprate.LimitAll(..).
	//
	// Please see _example/main.go for other more, or read the library code.
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	return r
}

func main() {
	loadEnv()

	lookupDb, err := loadDb()
	if err != nil {
		log.Fatalln("Error loading id db:", err)
	}

	lookupClient := lookup.New(lookupDb)

	r := setupRouter()
	r.Post("/lookup", lookupClient.LookupHandler)

	log.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
