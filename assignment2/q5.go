package main

import "fmt"

func q5() {
	fmt.Println("Welcome to Calculator")
	fmt.Println("enter two numbers")
	var a, b float32
	var op string
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("Enter the operator as (+,-,*,/)")
	fmt.Scan(&op)
	result := calculator(a, b, op)
	fmt.Println("The result is", result)
}
func calculator(a float32, b float32, op string) (result float32) {
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("Invalid")
		}
	}
	return
}
