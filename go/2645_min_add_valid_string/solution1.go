package main

func addMinimum1(word string) int {

	if word == "abc" {
		return 0
	}
	changes := 0

	var newWord []byte
	wordLen := len(word)
	for idx := 0; idx < wordLen; idx++ {
		if word[idx] == 'b' {
			idxN := idx + 1
			changes++ // a
			if idxN < wordLen && word[idxN] == 'c' {
				idx = idxN
			} else {
				changes++ // c
			}
		}

		if word[idx] == 'a' {
			idxN := idx + 1
			if idxN < wordLen && word[idxN] == 'b' {
				idx = idxN
				idxN += 1
				if idxN < wordLen && word[idxN] == 'c' {
					idx = idxN
				} else {
					changes++ // c
				}
			} else {
				changes += 2 // b,c
			}
		}
		newWord = append(newWord, []byte("abc")...) // build the final word; unuseds
	}
	return changes
}
