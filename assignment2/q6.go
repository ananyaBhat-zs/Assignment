package main

import "fmt"

func q6() {
	var n int
	fmt.Println("Enter a no")
	fmt.Scan(&n)
	s := sum(n)
	fmt.Println("Sum:", s)
}
func sum(n int) (s int) {
	for i := 0; i <= n; i++ {
		s += i
	}
	return
}
