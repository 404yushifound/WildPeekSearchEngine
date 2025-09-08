package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Animal struct {
	Name    string `json:"name"`
	Color   string `json:"color"`
	Height  string `json:"height"`
	Weight  string `json:"weight"`
	Species string `json:"species"`
}

func main() {
	// Step 1: Data of Exotic Animals
	animals := []Animal{
		{"Macaw", "Blue-Yellow", "85 cm", "1.2 kg", "Bird"},
		{"Axolotl", "Pink", "23 cm", "250 g", "Amphibian"},
		{"Fennec Fox", "Cream", "20 cm", "1.5 kg", "Mammal"},
		{"Red Panda", "Reddish-Brown", "60 cm", "6 kg", "Mammal"},
		{"Koi Fish", "Orange-White", "70 cm", "4 kg", "Fish"},
	}

	// Step 2: Index
	index := make(map[string][]int)
	for i, a := range animals {
		fields := []string{a.Name, a.Color, a.Height, a.Weight, a.Species}
		for _, field := range fields {
			words := strings.Fields(strings.ToLower(field))
			for _, word := range words {
				index[word] = append(index[word], i)
			}
		}
	}

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Exotic Animal Search API! Use /search?q=<term>")
	})

	// Search endpoint
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := strings.ToLower(r.URL.Query().Get("q"))
		if query == "" {
			http.Error(w, "Please provide a search term using ?q=...", http.StatusBadRequest)
			return
		}

		results, found := index[query]
		if !found {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[]`))
			return
		}

		var matched []Animal
		seen := make(map[int]bool)
		for _, idx := range results {
			if seen[idx] {
				continue
			}
			seen[idx] = true
			matched = append(matched, animals[idx])
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(matched)
	})

	// Start server
	port := "10000"
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}

