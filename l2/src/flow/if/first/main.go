package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		if i%2 == 0 {
			fmt.Print(" divisible by 2")
		} else if i%3 == 0 {
			fmt.Print(" divisible by 3")
		} else if i%5 == 0 {
			fmt.Print(" divisible by 5")
		}
		fmt.Println()
	}
}
