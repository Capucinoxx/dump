package main

import (
	"testing"
)

type mockRand struct {
	value int
}

func (m *mockRand) Intn(n int) int {
	return m.value
}

func TestGenerateWord(t *testing.T) {
	t.Run("La fonction devrait retourner un mot", func(t *testing.T) {
		wordle := &Wordle{}

		wordle.GenerateWord()

		if wordle.word == "" {
			t.Error("GenerateWord() = \"\"; Nous attendons un mot non vide")
		}
	})

	t.Run("La fonction devrait retourner un mot affecté par la fonction Intn", func(t *testing.T) {
		wordle := &Wordle{randomizer: &mockRand{value: 2}}

		wordle.GenerateWord()

		if wordle.word != "image" {
			t.Errorf("GenerateWord() = %s; Nous attendons image", wordle.word)
		}
	})
}

func TestTry(t *testing.T) {
	wordle := &Wordle{word: "golang"}

	tests := map[string]struct {
		guess       string
		wantClue    string
		wantCorrect bool
	}{
		"le mot est incorrect":                         {"python", "⬛⬛⬛⬛🟨🟨", false},
		"le mot est correct":                           {"golang", "🟩🟩🟩🟩🟩🟩", true},
		"le mot est partiellement correct":             {"google", "🟩🟩🟨🟨🟨⬛", false},
		"le mot est incorrect et de taille différente": {"wwwwwwwwwwwwww", "⬛⬛⬛⬛⬛⬛", false},
		"le mot est plus court que celui attendu":      {"go", "🟩🟩⬛⬛⬛⬛", false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			clue, correct := wordle.Try(tt.guess)
			if clue != tt.wantClue || correct != tt.wantCorrect {
				t.Errorf("Try(%s) = %s, %t; Nous voulons %s, %t", tt.guess, clue, correct, tt.wantClue, tt.wantCorrect)
			}
		})
	}
}
