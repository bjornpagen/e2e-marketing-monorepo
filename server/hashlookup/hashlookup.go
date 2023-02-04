package hashlookup

import (
	"encoding/json"
	"errors"
	"net/http"
)

type HashLookupClient struct {
	hashToEmailMap map[string]string
}

// New creates a new client for the hashlookup package.
func New(m map[string]string) *HashLookupClient {
	return &HashLookupClient{
		hashToEmailMap: m,
	}
}

type HashLookupRequest struct {
	Hash string `json:"hash"`
}

type HashLookupResponse struct {
	Email string `json:"email"`
}

func (c *HashLookupClient) HashLookupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req HashLookupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	hash := req.Hash

	email, err := c.Lookup(hash)
	if err != nil {
		http.Error(w, "email not found", http.StatusNotFound)
		return
	}

	resp := HashLookupResponse{
		Email: email,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

// Lookup takes a salted Blake3 hash of an email address and returns the email address if found, or an error if not found.
func (c *HashLookupClient) Lookup(hash string) (string, error) {
	email, ok := c.hashToEmailMap[hash]
	if !ok {
		return "", errors.New("email not found")
	}
	return email, nil
}
