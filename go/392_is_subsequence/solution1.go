package main

// REMEMBER:  s == subsequence (?), t == WORD
func isSubsequence(s string, t string) bool {
	strIdx := 0

	// iterate over WHOLE word (t); keep track of idx in subsequence (s)
	for i := 0; i < len(t) && strIdx < len(s); i++ {
		if t[i] == s[strIdx] {
			strIdx++
		}
	}
	return strIdx == len(s)
}
