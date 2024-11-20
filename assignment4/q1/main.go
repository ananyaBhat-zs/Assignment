package main

import (
	"assignment4.q1/stack"
	"fmt"
	"os"
)

func main() {
	s := stack.NewStack()
	for {
		fmt.Println("Enter if you wish to push(1) or pop(2) or exit(3)")
		ch := 0
		fmt.Scan(&ch)
		switch ch {
		case 1:
			{
				var val int
				fmt.Println("Enter val")
				fmt.Scan(&val)
				s.Push(val)
			}
		case 2:
			{
				if flag, val := s.Pop(); flag == true {
					fmt.Println("The number that was popped:", val)
				} else {
					fmt.Println("No items in the stack", val)
				}

			}
		case 3:
			{
				os.Exit(0)
			}
		}
	}

}
