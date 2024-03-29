package main

import (
	"math/rand"
	"strings"
	"time"
)

var dictionary = []string{
	"docker",
	"container",
	"image",
	"golang",
	"igl601",
}

type Randomizer interface {
	Intn(n int) int
}

type WordleGame interface {
	GenerateWord()
	Try(guess string) (string, bool)
	WordLength() int
}

type Wordle struct {
	word       string
	randomizer Randomizer
}

func (w *Wordle) GenerateWord() {
	if w.randomizer == nil {
		w.randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	w.word = dictionary[w.randomizer.Intn(len(dictionary))]
}

func (w *Wordle) Try(guess string) (string, bool) {
	clues := make([]string, len(w.word))

	for i := len(guess); i < len(w.word); i++ {
		guess += string(rune(0))
	}
	guess = guess[:len(w.word)]

	for i, g := range guess {
		if g == rune(w.word[i]) {
			clues[i] = "🟩"
		} else if strings.ContainsRune(w.word, g) {
			clues[i] = "🟨"
		} else {
			clues[i] = "⬛"
		}
	}

	return strings.Join(clues, ""), guess == w.word
}

func (w *Wordle) WordLength() int {
	return len(w.word)
}
