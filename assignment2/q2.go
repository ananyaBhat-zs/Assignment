package main

import "fmt"

func q2() {
	var a, b, c int
	fmt.Println("Enter three no's")
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Println(a, "is the largest number")
		if b > c {
			fmt.Println(b, "is the second largest number")
		} else {
			fmt.Println(c, "is the second largest number")
		}
	} else if b > a && b > c {
		fmt.Println(b, "is the largest number")
		if a > c {
			fmt.Println(b, "is the second largest number")
		} else {
			fmt.Println(c, "is the second largest number")
		}
	} else {
		fmt.Println(c, "is the largest number")
		if b > a {
			fmt.Println(b, "is the second largest number")
		} else {
			fmt.Println(a, "is the second largest number")
		}
	}
}
