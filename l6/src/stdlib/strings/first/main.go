package main

import (
	"fmt"
	"strings"
)

func main() {
	// START OMIT
	// true
	fmt.Println(strings.Contains("test", "es"))
	// 2
	fmt.Println(strings.Count("test", "t"))
	// true
	fmt.Println(strings.HasPrefix("test", "te"))
	// true
	fmt.Println(strings.HasSuffix("test", "st"))
	// 1
	fmt.Println(strings.Index("test", "e"))
	// "a-b"
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	// == "aaaaa"
	fmt.Println(strings.Repeat("a", 5))
	// "bbaa"
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	// []string{"a","b","c","d","e"}
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	// "test"
	fmt.Println(strings.ToLower("TEST"))
	// "TEST"
	fmt.Println(strings.ToUpper("test"))
	// END OMIT
}
