package main

import "testing"

func TestAddMinimum1(t *testing.T) {
	word := "b"
	expected := 2
	result := addMinimum1(word)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	word = "aaa"
	expected = 6
	result = addMinimum1(word)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	word = "abc"
	expected = 0
	result = addMinimum1(word)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
