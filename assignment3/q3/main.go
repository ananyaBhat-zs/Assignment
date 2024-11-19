package main

import "fmt"

func main() {
	var m = map[string][]int{
		"A": {1, 2, 3, 4, 5},
		"B": {6, 7, 8, 9},
		"C": {10, 11, 12, 13, 14, 15, 16},
		"D": {17, 18, 19, 20, 21},
	}
	newMap := sumValuesByKey(&m)
	fmt.Println(newMap)
}
func sumValuesByKey(m *map[string][]int) map[string]int {
	var newMap = make(map[string]int)
	for k, v := range *m {
		sum := 0
		for i := 0; i < len(v); i++ {
			sum += v[i]
		}
		newMap[k] = sum
	}
	return newMap
}
