package main

import "testing"

func Test2390_1(t *testing.T) {
	str := "leet**cod*e"
	expect := "lecoe"

	result := removeStars(str)
	if result != expect {
		t.Errorf("Expecting %v, got %v", expect, result)
	}
}

func Test2390_2(t *testing.T) {
	str := "erase*****"
	expect := ""

	result := removeStars(str)
	if result != expect {
		t.Errorf("Expecting %v, got %v", expect, result)
	}

}

func Test2390_3(t *testing.T) {
	str := "*"
	expect := ""

	result := removeStars(str)
	if result != expect {
		t.Errorf("Expecting %v, got %v", expect, result)
	}

}
