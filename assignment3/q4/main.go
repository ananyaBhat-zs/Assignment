package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	m := convertToMap(&s)
	fmt.Println(m)
}
func convertToMap(s *[]int) map[int]int {
	m := make(map[int]int)
	for i, v := range *s {
		m[i] = v
	}
	return m
}
