package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 0 {
			fmt.Println("Zero")
		} else if i == 1 {
			fmt.Println("One")
		} else if i == 2 {
			fmt.Println("Two")
		} else if i == 3 {
			fmt.Println("Three")
		} else if i == 4 {
			fmt.Println("Four")
		} else if i == 5 {
			fmt.Println("Five")
		} else {
			fmt.Println(i)
		}
	}
}
