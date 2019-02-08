// Package dict is a Toaq dictionary.
package dict

import (
	"encoding/json"
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
func Load(url string) (map[string]Word, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dict struct {
		Entries []Word `json:"entries"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&dict); err != nil {
		return nil, err
	}
	ws := make(map[string]Word)
	for i := range dict.Entries {
		e := dict.Entries[i]
		e.Word = strings.ToLower(tone.Strip(e.Word))
		ws[e.Word] = e
	}
	return ws, nil
}
