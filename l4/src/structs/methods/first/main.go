package main

import "fmt"

type Person struct {
	Name string
}

// START PTR OMIT
func (p *Person) SetNamePtr(name string) Person {
	p.Name = name
	return *p
}

// END PTR OMIT

// START OMIT
func (p Person) SetName(name string) Person {
	p.Name = name
	return p
}

// END OMIT

func main() {
	p := Person{Name: "Bob Kelso"}
	p1 := p.SetName("John Dorian")
	fmt.Println(p, p1)
	p2 := p.SetNamePtr("Perry Kox")
	fmt.Println(p, p2)
}
