package main

func addMinimum2(word string) int {
	changes := 0
	wanted := 'a'

	// iterate over the word runes
	for _, letter := range word {

		for letter != wanted {
			changes++ // track adds
			wanted++  // advance the character
			if wanted > 'c' {
				wanted = 'a'
			}
		}

		wanted++
		if wanted > 'c' {
			wanted = 'a'
		}
	}

	// anything missing after last character in word?
	if wanted == 'b' {
		changes += 2 // b,c
	} else if wanted == 'c' {
		changes += 1 // c
	}

	return changes
}
