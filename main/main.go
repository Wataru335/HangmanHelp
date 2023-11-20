package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/01-edu/z01"
)

func main() {

	filename := os.Args[1] // Récupération du nom du fichier passé en argument
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
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

	fmt.Println("\nMot à deviner :", string(guessWord))
	for _, line := range lines {
		fmt.Println(line)
	}

	z01.PrintRune('\n')
	fmt.Printf("Vous avez 10 tentatives pour deviner le mot.")
	z01.PrintRune('\n')

	for attempts > 0 {
		fmt.Println("--------------------------------------------------")
		fmt.Print("\nSuggérez une lettre ou devinez le mot complet : ")
		var guess string
		fmt.Scanln(&guess)

		if guessedLetters[guess] {
			fmt.Println("Vous avez déjà entré cette lettre ou ce mot. Veuillez choisir une autre lettre ou mot.")
			z01.PrintRune('\n')
			continue
		}

		guessedLetters[guess] = true

		if len(guess) == 1 {
			if strings.Contains(word, guess) {
				fmt.Println("\nLa lettre est présente dans le mot !")
				z01.PrintRune('\n')
				maj_mot(word, guess, guessWord)
			} else {
				fmt.Println("\nLa lettre n'est pas présente dans le mot.")
				z01.PrintRune('\n')
				attempts--
				displayHangman(10 - attempts)
				fmt.Printf("Il vous reste %d tentatives.\n", attempts)
				z01.PrintRune('\n')
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

			fmt.Println("\nLe mot complet n'est pas correct. Vous avez fait", wrongCount, "erreurs.")
			z01.PrintRune('\n')
			attempts -= wrongCount
			fmt.Printf("Il vous reste %d tentatives.\n", attempts)
			z01.PrintRune('\n')
		} else if len(guess) == len(word) && guess == word {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("\nFélicitations ! Vous avez deviné le mot : ", word)
			if attempts < 10 {
				fmt.Println("en :", 10-attempts, "tentative(s)")
				break
			} else {
				fmt.Println("Vous avez fait un score parfait !")
				break
			}
		}

		fmt.Println("Mot à deviner :", string(guessWord))
		z01.PrintRune('\n')

		if string(guessWord) == word { // Si le mot est trouvé
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("\nFélicitations ! Vous avez deviné le mot : ", word)
			if attempts < 10 {
				fmt.Println("en :", 10-attempts, "tentative")
				break
			} else {
				fmt.Println("Vous avez fait un score parfait !")
				break
			}
		}
	}

	if attempts == 0 { // Si le mot n'est pas trouvé
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("\nDésolé, vous avez épuisé toutes vos tentatives. Vous êtes désormais pendu(e). \nLe mot à deviner était :", word)
		displayHangman(10 - attempts)
	}
}

func maj_mot(word, guess string, guessWord []byte) { // Mise à jour du mot
	for i := 0; i < len(word); i++ {
		if word[i] == guess[0] {
			guessWord[i] = guess[0]
		}
	}
}

var lines []string

func readLinesFromFile() ([]string, error) { // Lecture du fichier hangman.txt
	file, err := os.Open("hangman.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func displayHangman(currentAttempt int) { // Affichage du pendu
	readLinesFromFile()

	n := len(lines) / 8
	if currentAttempt >= 1 && currentAttempt < n {
		startIndex := currentAttempt * 8
		endIndex := (currentAttempt + 1) * 8
		if endIndex > len(lines) {
			endIndex = len(lines)
		}

		for i := startIndex; i < endIndex; i++ {
			fmt.Println(lines[i])
		}
	} else {
		fmt.Println("Invalid value for 'currentAttempt'.")
	}
	return
}
