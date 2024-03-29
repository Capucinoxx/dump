package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type ApiResponse struct {
	Guess    string `json:"guess"`
	Clues    string `json:"clues"`
	Finished bool   `json:"finished"`
}

var (
	mu   sync.Mutex
	game WordleGame
)

func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/new", logging(newGameHandler))
	http.HandleFunc("/guess", logging(guessHandler))
	http.HandleFunc("/length", logging(lengthHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	game = &Wordle{}
	game.GenerateWord()

	w.WriteHeader(http.StatusOK)
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Guess string `json:"guess"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la requête", http.StatusBadRequest)
		return
	}

	clues, finished := game.Try(request.Guess)

	response := ApiResponse{
		Guess:    request.Guess,
		Clues:    clues,
		Finished: finished,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func lengthHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	if game == nil {
		game = &Wordle{}
		game.GenerateWord()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"length": game.WordLength()})
}

func logging(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requête reçue: %s %s à %s", r.Method, r.URL.Path, time.Now().Format(time.RFC1123))
		nextFunc(w, r)
	}
}
