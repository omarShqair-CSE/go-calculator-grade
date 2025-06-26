// Write a program that returns the grade based on a numeric score.

// Implement a simple calculator that asks for two numbers and an operator (+ - * /).
package main

import (
	"fmt"
)

func grade(score int) {
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}

func calculator(x, y int, op rune) (int, error) {
	switch op {
	case '+':
		return x + y, nil
	case '-':
		return x - y, nil
	case '*':
		return x * y, nil
	case '/':
		if y == 0 {
			return 0, fmt.Errorf("can't divide by zero")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("unknown operator: %c", op)
	}
}

func main() {
	var score int
	fmt.Print("Enter your grade (0-100): ")
	_, err := fmt.Scan(&score)
	if err != nil {
		fmt.Println("INVALID number")
		return
	}
	grade(score)

	var n1, n2 int
	fmt.Print("Enter number 1: ")
	_, err = fmt.Scan(&n1)
	if err != nil {
		fmt.Println("INVALID number")
		return
	}

	fmt.Print("Enter number 2: ")
	_, err = fmt.Scan(&n2)
	if err != nil {
		fmt.Println("INVALID number")
		return
	}

	var operator string
	fmt.Print("Enter operator (+ - * /): ")
	_, err = fmt.Scan(&operator)
	if err != nil || len(operator) == 0 {
		fmt.Println("INVALID operator")
		return
	}
	op := rune(operator[0])

	result, err := calculator(n1, n2, op)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %d\n", result)
}
