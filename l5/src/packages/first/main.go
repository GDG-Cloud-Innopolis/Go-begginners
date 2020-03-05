package main

import (
	"fmt"

	"github.com/GDG-Cloud-Innopolis/Go-begginners/l5/src/packages/first/math" // HL01
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
