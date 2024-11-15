package main

import "fmt"

func q1() {
	var marks int
	fmt.Println("Enter your marks")
	fmt.Scan(&marks)
	if marks > 90 && marks < 100 {
		fmt.Println("Your grade is Outstanding:O")
	} else if marks > 80 && marks < 90 {
		fmt.Println("Your grade is Excellent: A")
	} else if marks > 70 && marks < 80 {
		fmt.Println("Your grade is Very Good:B")
	} else if marks > 60 && marks < 70 {
		fmt.Println("Your grade is Good:C")
	} else {
		fmt.Println("You can work on your score")
	}
}
