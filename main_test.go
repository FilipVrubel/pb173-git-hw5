package main

import "testing"

func TestAddPositiveNumbers(t *testing.T) {
	result := add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("add(3, 5) = %d; want %d", result, expected)
	}
}

func TestAddNegativeNumbers(t *testing.T) {
	result := add(-4, -6)
	expected := -10
	if result != expected {
		t.Errorf("add(-4, -6) = %d; want %d", result, expected)
	}
}