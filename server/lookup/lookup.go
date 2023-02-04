package lookup

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Id string
type User struct {
	Email string `json:"email"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type LookupClient struct {
	db map[Id]User
}

// New creates a new client for the idlookup package.
func New(db map[Id]User) *LookupClient {
	return &LookupClient{
		db: db,
	}
}

type Request struct {
	Id Id `json:"id"`
}
type Response User

func (c *LookupClient) LookupHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	id := Id(req.Id)
	log.Printf("Looking up id %s", id)

	user, err := c.lookup(id)
	if err != nil {
		http.Error(w, "Error looking up id", http.StatusBadRequest)
		return
	}

	resp := Response(user)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(resp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// lookup takes an id of an email address and returns the email address if found, or an error if not found.
func (c *LookupClient) lookup(id Id) (User, error) {
	user, ok := c.db[id]
	if !ok {
		return User{}, errors.New("id not found")
	}

	return user, nil
}
