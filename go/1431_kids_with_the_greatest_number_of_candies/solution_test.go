package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	expected := []bool{true, true, true, false, true}
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	expected = []bool{true, false, false, false, false}
	result = kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	expected = []bool{true, false, true}
	result = kidsWithCandies([]int{12, 1, 12}, 10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
