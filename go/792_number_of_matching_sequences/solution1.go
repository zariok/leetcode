package main

func numMatchingSubseq(s string, words []string) int {
	valid := 0
	for _, word := range words {
		if isSubsequence(word, s) {
			valid++
		}
	}
	return valid
}

func isSubsequence(sub string, word string) bool {
	subIdx := 0

	for i := 0; i < len(word) && subIdx < len(sub); i++ {
		if word[i] == sub[subIdx] {
			subIdx++
		}
	}

	return subIdx == len(sub)
}
