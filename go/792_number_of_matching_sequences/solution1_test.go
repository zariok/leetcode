package main

import "testing"

func Test792(t *testing.T) {
	expect := 3
	result := numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"})
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	expect = 2
	result = numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"})
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

}
