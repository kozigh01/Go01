package interfaces

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type Person2 struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func interfaces10() {
	a := Person{"Arthur Dent", 42}
	b := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println(a, b)

	var c fmt.Stringer = Person{"Me", 55}
	d := Person2{"Dude", 24}
	fmt.Println(c)
	fmt.Println(d)
}