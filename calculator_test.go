package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 2)
	expected := 3
	if result != float64(expected) {
		t.Errorf("Add(1, 2) = %f; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 2)
	expected := 3
	if result != float64(expected) {
		t.Errorf("Subtract(5, 2) = %f; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(3, 4)
	expected := 12
	if result != float64(expected) {
		t.Errorf("Multiply(3, 4) = %f; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := divide(10, 2)
	expected := 5
	if result != float64(expected) {
		t.Errorf("Divide(10, 2) = %f; want %d", result, expected)
	}
}
