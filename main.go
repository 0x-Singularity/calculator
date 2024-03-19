package main

import (
  "fmt"
  "os"
  "strconv"
  "github.com/0x-Singularity/calculator"
)

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
    result = calculator.Add(operand1, operand2)
  case "subtract":
    result = calculator.Subtract(operand1, operand2)
  case "multiply":
    result = calculator.Multiply(operand1, operand2)
  case "divide":
    result = calculator.Divide(operand1, operand2)
  default:
    fmt.Println("Error: Invalid operation")
    os.Exit(1)
  }

  fmt.Println("Result:", result)
}

