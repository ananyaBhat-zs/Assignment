package main

import "fmt"

func q4() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)
	if isEven(n) {
		fmt.Println("The no", n, "is even")
	} else {
		fmt.Println("The no", n, "is odd")
	}
}
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
