package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	expected := []int{0, 1}
	result := twoSum2([]int{2, 7, 11, 15}, 9)
	slices.Sort(result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	expected = []int{1, 2}
	result = twoSum2([]int{3, 2, 4}, 6)
	slices.Sort(result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSum2([]int{3, 3}, 6)
	slices.Sort(result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
