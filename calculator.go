package calculator

import (
	"fmt"
	"os"
  "strconv"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func ParseArgs() (string, float64, float64, error) {
	if len(os.Args) < 4 {
		return "", 0, 0, fmt.Errorf("Usage: calculator <operation> <operand1> <operand2>")
	}
	operation := os.Args[1]
	operand1, err1 := strconv.ParseFloat(os.Args[2], 64)
	operand2, err2 := strconv.ParseFloat(os.Args[3], 64)
	if err1 != nil || err2 != nil {
		return "", 0, 0, fmt.Errorf("Error: Invalid operands")
	}
	return operation, operand1, operand2, nil
}


