package main

import "testing"

func Test104(t *testing.T) {
	root := []interface{}{3, 9, 20, nil, nil, 15, 7}
	expected := 3
	if got := maxDepth(buildTree(root)); got != expected {
		t.Errorf("maxDepth() = %v, want %v", got, expected)
	}

	root = []interface{}{1, nil, 2}
	expected = 2
	if got := maxDepth(buildTree(root)); got != expected {
		t.Errorf("maxDepth() = %v, want %v", got, expected)
	}
}
