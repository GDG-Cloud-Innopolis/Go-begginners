package main

import "fmt"

// START PERSON OMIT
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// END PERSON OMIT

// START ANDROID OMIT
type Android struct {
	Person
	Model string
}

// END ANDROID OMIT

func main() {
	a := new(Android)
	a.Talk()
}
