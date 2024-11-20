package main

import (
	"assignment4.q2.a/calculator"
	"fmt"
)

func main() {
	var v1, v2 float64
	var ch string
	fmt.Println("Enter the two numbers")
	fmt.Scan(&v1, &v2)
	c := calculator.NewCalc(v1, v2)
	fmt.Println("Enter the operation")
	fmt.Scan(&ch)
	switch ch {
	case "+":
		{
			fmt.Printf("The sum of %v and %v is %v", v1, v2, c.Add())
		}
	case "-":
		{
			fmt.Printf("The difference of %v and %v is %v", v1, v2, c.Sub())
		}
	case "*":
		{
			fmt.Printf("The product of %v and %v is %v", v1, v2, c.Mul())
		}
	case "/":
		{
			if q, flag := c.Div(); flag == true {
				fmt.Printf("The quotient of %v and %v is %v", v1, v2, q)
			} else {
				fmt.Println("Division by zero is not possible")
			}

		}
	}
}
