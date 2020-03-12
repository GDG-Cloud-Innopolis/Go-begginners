package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Cannot create file")
	}
	defer file.Close()

	file.WriteString("test")
}
