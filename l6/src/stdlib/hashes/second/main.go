package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

// START FIRST OMIT
func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

// END FIRST OMIT

// START SECOND OMIT
func main() {
	fmt.Println("Hello")
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}

// END SECOND OMIT
