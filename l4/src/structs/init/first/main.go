package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{}
	fmt.Printf("%v\n", c)
}
