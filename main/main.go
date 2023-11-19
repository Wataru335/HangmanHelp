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

	revealedIndices := make([]int, 2) // Révélation de 2 lettres du mot
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
				fmt.Printf("Il vous reste %d tentatives.\n", attempts)
				z01.PrintRune('\n')
			}
		} else if len(guess) == len(word) && guess != word { // Si le mot est faux
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
				fmt.Println("en :", 10-attempts, "tentative")
				break
			} else {
				fmt.Println("Vous avez fait un score parfait !")
				break
			}
		}

		fmt.Println("Mot à deviner :", string(guessWord))
		z01.PrintRune('\n')
		affichage_Hangman(attempts)

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
		affichage_Hangman(attempts)
	}
}

func maj_mot(word, guess string, guessWord []byte) { // Mise à jour du mot
	for i := 0; i < len(word); i++ {
		if word[i] == guess[0] {
			guessWord[i] = guess[0]
		}
	}
}

func affichage_Hangman(attempts int) { // Affichage du pendu
	hangman0 := []string{
		"\n",
	}
	hangman1 := []string{
		"\n\n\n\n\n\n\n\n ========= \n\n",
	}
	hangman2 := []string{
		"\n       |\n       |\n       |\n       |\n       |\n ========= \n\n",
	}
	hangman3 := []string{
		"\n   +---+\n       |\n       |\n       |\n       |\n       |\n ========= \n\n",
	}
	hangman4 := []string{
		"\n   +---+\n   |   |\n       |\n       |\n       |\n       |\n ========= \n\n",
	}
	hangman5 := []string{
		"\n   +---+\n   |   |\n   O   |\n       |\n       |\n       |\n ========= \n\n",
	}
	hangman6 := []string{
		"\n   +---+\n   |   |\n   O   |\n   |   |\n       |\n       |\n ========= \n\n",
	}
	hangman7 := []string{
		"\n   +---+\n   |   |\n   O   |\n  /|   |\n       |\n       |\n ========= \n\n",
	}
	hangman8 := []string{
		"\n   +---+\n   |   |\n   O   |\n  /|\\  |\n       |\n       |\n ========= \n\n",
	}
	hangman9 := []string{
		"\n   +---+\n   |   |\n   O   |\n  /|\\  |\n  /    |\n       |\n ========= \n\n",
	}
	hangman10 := []string{
		"\n   +---+\n   |   |\n   O   |\n  /|\\  |\n  / \\  |\n       |\n ========= \n\n",
	}

	hangman := [][]string{hangman0, hangman1, hangman2, hangman3, hangman4, hangman5, hangman6, hangman7, hangman8, hangman9, hangman10}

	concatenatedHangman0 := strings.Join(hangman[0], "")
	concatenatedHangman1 := strings.Join(hangman[1], "")
	concatenatedHangman2 := strings.Join(hangman[2], "")
	concatenatedHangman3 := strings.Join(hangman[3], "")
	concatenatedHangman4 := strings.Join(hangman[4], "")
	concatenatedHangman5 := strings.Join(hangman[5], "")
	concatenatedHangman6 := strings.Join(hangman[6], "")
	concatenatedHangman7 := strings.Join(hangman[7], "")
	concatenatedHangman8 := strings.Join(hangman[8], "")
	concatenatedHangman9 := strings.Join(hangman[9], "")
	concatenatedHangman10 := strings.Join(hangman[10], "")

	if attempts < 0 {
		attempts = 0
	}
	if attempts >= len(hangman) {
		attempts = len(hangman) - 1
	}
	if attempts == 10 {
		fmt.Println(concatenatedHangman0)
	}
	if attempts == 9 {
		fmt.Println(concatenatedHangman1)
	}
	if attempts == 8 {
		fmt.Println(concatenatedHangman2)
	}
	if attempts == 7 {
		fmt.Println(concatenatedHangman3)
	}
	if attempts == 6 {
		fmt.Println(concatenatedHangman4)
	}
	if attempts == 5 {
		fmt.Println(concatenatedHangman5)
	}
	if attempts == 4 {
		fmt.Println(concatenatedHangman6)
	}
	if attempts == 3 {
		fmt.Println(concatenatedHangman7)
	}
	if attempts == 2 {
		fmt.Println(concatenatedHangman8)
	}
	if attempts == 1 {
		fmt.Println(concatenatedHangman9)
	}
	if attempts == 0 {
		fmt.Println(concatenatedHangman10)
	}
}
