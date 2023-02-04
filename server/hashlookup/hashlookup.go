package hashlookup

import (
	"encoding/json"
	"net/http"
)

var hashMap = map[string]string{
	"<hash1>": "user1@example.com",
	"<hash2>": "user2@example.com",
	// ... Add more hashes and email addresses
}

// HashLookupHandler is a http.Handler that handles incoming HTTP requests
// and returns the email address corresponding to the provided hash.
func HashLookupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Hash string `json:"hash"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email, ok := hashMap[request.Hash]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(struct {
		Email string `json:"email"`
	}{
		Email: email,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
