package main

import (
	"testing"
)

func TestGCDOfStrings2(t *testing.T) {
	str1 := "ABCABC"
	str2 := "ABC"
	expected := "ABC"
	result := gcdOfStrings2(str1, str2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	str1 = "ABABAB"
	str2 = "ABAB"
	expected = "AB"
	result = gcdOfStrings2(str1, str2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	str1 = "LEET"
	str2 = "CODE"
	expected = ""
	result = gcdOfStrings2(str1, str2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	str1 = "ABCDEF"
	str2 = "ABC"
	expected = ""
	result = gcdOfStrings2(str1, str2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	str1 = "ABABABAB"
	str2 = "ABAB"
	expected = "ABAB"
	result = gcdOfStrings2(str1, str2)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
