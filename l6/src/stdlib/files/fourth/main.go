package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Cannot open directory")
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Cannot read directory")
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
	// END OMIT
}
