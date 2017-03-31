package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	n := 9544591

	expected := true
	actual := isPalindrome(n)

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
