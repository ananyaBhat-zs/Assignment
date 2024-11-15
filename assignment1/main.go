package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var name string
	var r, s float64
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	greet(name)
	fmt.Println("Enter a number")
	fmt.Scan(&n)
	fmt.Println("Entered no is:", n)
	var twice = calculateDouble(n)
	fmt.Println("The double of the no", n, "is", twice)
	var doubleDataTypeVal = convertToDouble(n)
	fmt.Printf("The double data type value of %v is %v and its data type is %T\n", n, doubleDataTypeVal, doubleDataTypeVal)
	fmt.Println("Enter radius of the circle")
	fmt.Scan(&r)
	var pc = perimeterC(r)
	fmt.Println("The perimeter of the circle with radius", r, "is", pc)
	fmt.Println("Enter the length of a side of the square")
	fmt.Scan(&s)
	var ps = perimeterS(s)
	fmt.Println("The perimeter of the square with side", s, "is", ps)
}
func calculateDouble(n int) int {
	return 2 * n
}
func convertToDouble(n int) float64 {
	return float64(n)
}
func greet(name string) {
	fmt.Println("Hello", name, "!")
}
func perimeterC(r float64) (perimeter float64) {
	perimeter = 2 * math.Pi * r
	return
}
func perimeterS(s float64) (perimeter float64) {
	perimeter = 4 * s
	return
}
