package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a, b float64) float64 {
	//Implement add function
  return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Cannot divide by zero")
		os.Exit(1)
	}
	return a / b
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
		os.Exit(1)
	}

	operation := os.Args[1]
	operand1, err1 := strconv.ParseFloat(os.Args[2], 64)
	operand2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid operands")
		os.Exit(1)
	}

	var result float64
	switch operation {
	case "add":
		result = add(operand1, operand2)
	case "subtract":
		result = subtract(operand1, operand2)
	case "multiply":
		result = multiply(operand1, operand2)
	case "divide":
		result = divide(operand1, operand2)
	default:
		fmt.Println("Error: Invalid operation")
		os.Exit(1)
	}

	fmt.Println("Result:", result)
}

