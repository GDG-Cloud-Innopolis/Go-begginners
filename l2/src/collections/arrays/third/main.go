package main

import "fmt"

func main() {
	var x = [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	for i, element := range x {
		fmt.Println(i)
		total += element
	}
	fmt.Println(total / float64(len(x)))
}
