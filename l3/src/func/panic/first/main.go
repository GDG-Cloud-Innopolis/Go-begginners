package main

import "fmt"

func main() {
	panic("PANIC")
	str := recover()
	fmt.Println(str)
}
