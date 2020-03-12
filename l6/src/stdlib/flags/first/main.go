package main

import (
	"flag"
	"fmt"
)

func main() {
	var minp int
	// Определение флагов
	maxp := flag.Int("max", 6, "the max value")
	flag.IntVar(&minp, "min", 1, "the min value")
	// Парсинг
	flag.Parse()
	fmt.Printf("min: %d, max: %d\n", minp, *maxp)
}
