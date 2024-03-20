package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OriginalQuote struct {
	Q string `json:"q"`
	A string `json:"a"`
}

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func getQuote() Quote {
	url := "https://zenquotes.io/api/random"

	response, err := http.Get(url)
	handleError(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		handleError(err)
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	handleError(err)

	var originalQuote []OriginalQuote
	err = json.Unmarshal(body, &originalQuote)
	handleError(err)

	return Quote{Quote: originalQuote[0].Q, Author: originalQuote[0].A}
}

func handleQuote(w http.ResponseWriter, r *http.Request) {
	quote := getQuote()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(quote)
	handleError(err)
}
