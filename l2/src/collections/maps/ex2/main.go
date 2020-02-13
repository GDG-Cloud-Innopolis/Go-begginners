package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	name, ok := elements["Un"]
	fmt.Println(name, ok)
}
