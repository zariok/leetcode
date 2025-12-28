package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	expected := true
	result := isSubsequence("abc", "ahbgdc")
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	expected = false
	result = isSubsequence("axc", "ahbgdc")
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
