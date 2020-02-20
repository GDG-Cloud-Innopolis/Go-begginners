package main

import "fmt"

func main() {
	a := 1
	pa := &a
	*pa = 2

	fmt.Printf("a = %v\n", a)
	fmt.Printf("data at %v is %v\n", pa, *pa)
}
