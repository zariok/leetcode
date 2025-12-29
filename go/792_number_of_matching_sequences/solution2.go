package main

// word and current character
type wordState struct {
	word string
	pos  int
}

func numMatchingSubseq2(s string, words []string) int {
	// Map each character to a list of words waiting for that character
	// Key: the next character each word is waiting for
	// Value: slice of wordState objects waiting for that character
	waiting := make(map[byte][]wordState)

	// Initialize: all words are waiting for their first character
	for _, word := range words {
		if len(word) > 0 {
			firstChar := word[0]
			waiting[firstChar] = append(waiting[firstChar], wordState{word: word, pos: 0})
		}
	}

	validCount := 0

	// Iterate through s once - this is the key efficiency gain
	for i := 0; i < len(s); i++ {
		char := s[i]

		// Get all words waiting for this character
		currentWaiting := waiting[char]
		delete(waiting, char)

		for _, ws := range currentWaiting {
			// Move to the next character in the word
			ws.pos++

			if ws.pos == len(ws.word) {
				validCount++
			} else {
				// continue working word; add it back to waiting list for its next character
				nextChar := ws.word[ws.pos]
				waiting[nextChar] = append(waiting[nextChar], ws)
			}
		}
	}

	return validCount
}
