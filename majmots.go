package hangmanhelp

func Maj_mot(word, guess string, guessWord []byte) { // Mise à jour du mot
	for i := 0; i < len(word); i++ {
		if word[i] == guess[0] {
			guessWord[i] = guess[0]
		}
	}
}