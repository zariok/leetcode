package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	expected := []int{0, 1}
	result := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	expected = []int{1, 2}
	result = twoSum([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSum([]int{3, 3}, 6)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
