package main

import "fmt"

func main() {
	x := [6]float64{
		98,
		93,
		77,
		82,
		83,
		88,
	}
	var total float64 = 0
	for _, element := range x {
		total += element
	}
	fmt.Println(total / float64(len(x)))
}
