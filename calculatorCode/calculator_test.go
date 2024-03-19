package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	if result != float64(expected) {
		t.Errorf("Add(1, 2) = %f; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	expected := 3
	if result != float64(expected) {
		t.Errorf("Subtract(5, 2) = %f; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	expected := 12
	if result != float64(expected) {
		t.Errorf("Multiply(3, 4) = %f; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	expected := 5.0 // Use float64 since result is a float64
	if result != expected {
		t.Errorf("Divide(10, 2) = %f; want %f", result, expected)
	}
	if err != nil {
		t.Errorf("Divide(10, 2) returned an error: %v", err)
	}

	// Test divide by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return an error")
	}


}
