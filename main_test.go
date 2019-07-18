package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(2,3)
	if total != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestSubstraction(t *testing.T) {
	total := subtraction(10, 5)
	if total != 5 {
		t.Errorf("Substraction was incorrect, got: %d, want: %d.", total, 5)
	}
}
