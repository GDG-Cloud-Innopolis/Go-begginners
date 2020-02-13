package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch {
		case i == 0:
			fmt.Println("Zero")
		case i == 1:
			fmt.Println("One")
		case i == 2:
			fmt.Println("Two")
		case i == 3:
			fmt.Println("Three")
		case i == 4:
			fmt.Println("Four")
		case i == 5:
			fmt.Println("Five")
		default:
			fmt.Println(i)
		}
	}
}
