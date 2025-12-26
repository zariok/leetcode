package main

import (
	"reflect"
	"testing"
)

func Test2215_1(t *testing.T) {
	int1 := []int{1, 2, 3}
	int2 := []int{2, 4, 6}

	expect := [][]int{{1, 3}, {4, 6}}

	pending := findDifference(int1, int2)

	// compare expect vs pending
	if !reflect.DeepEqual(pending, expect) {
		t.Errorf("Expected %v, got %v", expect, pending)
	}
}

func Test2215_2(t *testing.T) {
	int1 := []int{1, 2, 3, 3}
	int2 := []int{1, 1, 2, 2}

	expect := [][]int{{3}, {}}

	pending := findDifference(int1, int2)

	// compare expect vs pending
	if !reflect.DeepEqual(pending, expect) {
		t.Errorf("Expected %v, got %v", expect, pending)
	}
}
