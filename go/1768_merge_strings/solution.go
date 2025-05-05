package main

func mergeAlternately(word1 string, word2 string) string {
	var combined = ""
	w1len := len(word1)
	w2len := len(word2)
	maxlen := w1len
	if maxlen < w2len {
		maxlen = w2len
	}
	for x := 0; x < maxlen; x++ {
		if x < w1len {
			combined += string(word1[x])
		}
		if x < w2len {
			combined += string(word2[x])
		}
	}

	return combined
}
