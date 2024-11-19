package main

import "fmt"

func q7() {
	var x, y int
	fmt.Println("enter two no's")
	fmt.Scan(&x, &y)
	fmt.Println("Value of x is", x)
	fmt.Println("Value of y is", y)
	swap(&x, &y)
	fmt.Println("Value of x is", x)
	fmt.Println("Value of y is", y)
}
func swap(x, y *int) {
	*x, *y = *y, *x
}
