package pet

import (
	"encoding/json"
	"net/http"
)

type Pet struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	PhotoURLs []string `json:"photoUrls,omitempty"`
	Status    string   `json:"status,omitempty"`
	Tags      []Tag    `json:"tags,omitempty"`
}

type Tag struct {
	ID   int64
	Name string
}

// Handler returns pets
func Handler(w http.ResponseWriter, _ *http.Request) {
	var pets = []Pet{
		{ID: 1, Name: "Dog", PhotoURLs: []string{}, Status: "available", Tags: nil},
		{ID: 2, Name: "Cat", PhotoURLs: []string{}, Status: "pending", Tags: nil},
	}
	responseBytes, err := json.Marshal(&pets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = w.Write(responseBytes)
}
