package hangmanhelp

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Game() {

	file, err := os.Open("words.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano()) // Initialisation du générateur de mots aléatoires
	randomIndex := rand.Intn(len(words))
	word := words[randomIndex]
	guessWord := make([]byte, len(word))
	for i := range guessWord {
		guessWord[i] = '_'
	}
	revealedIndices := make([]int, 2) // Révélation des indices
	for i := 0; i < 2; i++ {
		unique := false
		for !unique {
			index := rand.Intn(len(word))
			unique = true
			for _, revealedIndex := range revealedIndices {
				if index == revealedIndex {
					unique = false
					break
				}
			}
			if unique {
				revealedIndices[i] = index
				guessWord[index] = word[index]
			}
		}
		for i, letter := range word {
			for _, revealedIndex := range revealedIndices {
				if i != revealedIndex && letter == rune(guessWord[revealedIndex]) {
					guessWord[i] = word[i]
				}
			}
		}
	}

	guessedLetters := make(map[string]bool)
	attempts := 10 // Nombre de tentatives

	for attempts > 0 {
		var guess string
		fmt.Scanln(&guess)

		if guessedLetters[guess] {
			continue
		}

		guessedLetters[guess] = true

		if len(guess) == 1 {
			if strings.Contains(word, guess) {
				Maj_mot(word, guess, guessWord)
			} else {
				attempts--
				DisplayHangman(10 - attempts)
			}
		} else if len(guess) == len(word) && guess != word { // Si le mot est incorrect
			var wrongCount int
			for i := 0; i < len(word); i++ {
				if word[i] == guess[i] {
					guessWord[i] = guess[i]
				} else {
					wrongCount++
				}
			}
			attempts -= wrongCount
		}
	}
}
