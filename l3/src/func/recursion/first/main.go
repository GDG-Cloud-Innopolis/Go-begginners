package main

import "fmt"

func factorial(x uint) uint {
	fmt.Println("Factorial called with ", x)
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func main() {
	fmt.Println("Factorial for 5 is ")
	fmt.Println(factorial(5))
}
