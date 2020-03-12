package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Cannot open file")
	}
	defer file.Close()

	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Cannot get file size")
	}
	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Cannot read file")
	}

	str := string(bs)
	fmt.Println(str)
	// END OMIT
}
