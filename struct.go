package main

import "fmt"

func main() {
	me := Person{"jaga", 26}
	me.introduce()
}

// type definition
type Person struct {
	name string
	age  int
}

// Person::introduce
func (p Person) introduce() {
	fmt.Printf("My name is %s, and I'm %d", p.name, p.age)
}
