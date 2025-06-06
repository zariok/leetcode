package main

import "testing"

func TestMinRemoveToMakeValid2(t *testing.T) {
	str := "lee(t(c)o)de)"
	expected := "lee(t(c)o)de"
	result := minRemoveToMakeValid2(str)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	str = "a)b(c)d"
	expected = "ab(c)d"
	result = minRemoveToMakeValid2(str)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	str = "))(("
	expected = ""
	result = minRemoveToMakeValid2(str)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	str = "())()((("
	expected = "()()"
	result = minRemoveToMakeValid2(str)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
