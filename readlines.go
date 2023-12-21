package hangmanhelp

import (
	"bufio"
	"os"
)

var lines []string

func ReadLinesFromFile() ([]string, error) { // Lecture du fichier hangman.txt
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
