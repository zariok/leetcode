package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	expected := "apbqcr"
	result := mergeAlternately(word1, word2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	word1 = "ab"
	word2 = "pqrs"
	expected = "apbqrs"
	result = mergeAlternately(word1, word2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	word1 = "abcd"
	word2 = "pq"
	expected = "apbqcd"
	result = mergeAlternately(word1, word2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
