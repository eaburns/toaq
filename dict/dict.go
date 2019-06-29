// Package dict is a Toaq dictionary.
package dict

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eaburns/toaq/tone"
)

//go:generate go run main.go -o words.go

// A Word contains all information about a single Toaq word.
type Word struct {
	// Word is the word in lower-case ASCII with no tone markers.
	Word string `json:"head"`
	// Def is the definition of the word.
	Def string `json:"body"`
	// By is the author or source of the word.
	By string `json:"by"`
	// Score is the score of the word in the toadua dictionary.
	Score int `json:"score"`
}

// Load loads the words from a URL serving them as a list of JSON.
func Load(url string, query string) (map[string]Word, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Success bool    `json:"success"`
		Entries *[]Word `json:"data"`
		Error   *string `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, fmt.Errorf("Toadua returned error: %s", *response.Error)
	}
	entries := *response.Entries
	ws := make(map[string]Word)
	for _, e := range entries {
		e.Word = strings.ToLower(tone.Strip(e.Word))
		ws[e.Word] = e
	}
	return ws, nil
}
