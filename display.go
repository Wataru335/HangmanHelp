package hangmanhelp

import "fmt"

func DisplayHangman(currentAttempt int) { // Affichage du pendu
	ReadLinesFromFile()

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
