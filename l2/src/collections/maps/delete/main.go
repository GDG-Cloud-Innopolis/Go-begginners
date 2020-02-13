package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println(elements)
	delete(elements, "He")
	fmt.Println(elements)
}
