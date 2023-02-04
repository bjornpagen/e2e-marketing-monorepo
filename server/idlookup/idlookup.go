package idlookup

import (
	"encoding/json"
	"errors"
	"net/http"
)

type IdLookupClient struct {
	idToEmailMap map[string]string
}

// New creates a new client for the idlookup package.
func New(m map[string]string) *IdLookupClient {
	return &IdLookupClient{
		idToEmailMap: m,
	}
}

type IdLookupRequest struct {
	Id string `json:"id"`
}

type IdLookupResponse struct {
	Email string `json:"email"`
}

func (c *IdLookupClient) IdLookupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req IdLookupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	email, err := c.Lookup(req.Id)
	if err != nil {
		http.Error(w, "email not found", http.StatusNotFound)
		return
	}

	resp := IdLookupResponse{
		Email: email,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

// Lookup takes an id of an email address and returns the email address if found, or an error if not found.
func (c *IdLookupClient) Lookup(id string) (string, error) {
	email, ok := c.idToEmailMap[id]
	if !ok {
		return "", errors.New("email not found")
	}
	return email, nil
}
