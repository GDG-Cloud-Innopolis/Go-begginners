package main

import "fmt"

func main() {
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr, str)
}
