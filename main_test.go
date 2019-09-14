package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 2 {
		t.Error("Test failed")
	}
}
