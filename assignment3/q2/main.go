package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var m = make(map[string]int)
	fmt.Println("Enter the word")
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		m[strings.ToLower(string(str[i]))] += 1
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
