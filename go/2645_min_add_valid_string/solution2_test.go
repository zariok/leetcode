package main

import "testing"

func TestAddMinimum2(t *testing.T) {
	word := "b"
	expected := 2
	result := addMinimum2(word)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	word = "aaa"
	expected = 6
	result = addMinimum2(word)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	word = "abc"
	expected = 0
	result = addMinimum2(word)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
