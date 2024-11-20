package main

import (
	"assignment4.q2.b/memorySupportCalculator"
	"fmt"
	"os"
)

func main() {
	ch := 0
	c := memorySupportCalculator.MemoryCalc{}
	for {
		fmt.Println("Do you want to perform SetMemory(1),AddToMemory(2),subtractFromMemory(3) or resetMemory(4) or exit(5)")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			{
				var n float64
				fmt.Println("Enter the no")
				fmt.Scan(&n)
				c.SetMemory(n)
			}
		case 2:
			{
				var n float64
				fmt.Println("Enter the no to be added")
				fmt.Scan(&n)
				fmt.Println("The sum is ", c.AddToMemory(n))
			}
		case 3:
			{
				var n float64
				fmt.Println("Enter the no to be subtracted from memory")
				fmt.Scan(&n)
				fmt.Println("The difference is", c.SubtractFromMemory(n))
			}
		case 4:
			{
				c.ResetMemory()
				fmt.Println("Memory reset to zero")
			}
		case 5:
			{
				os.Exit(0)
			}
		}
	}
}
