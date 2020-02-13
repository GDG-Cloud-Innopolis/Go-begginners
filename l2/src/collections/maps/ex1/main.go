package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"

	fmt.Println(elements["Un"])
}
