package functions

import (
	"fmt"
)

func Functions() {
	fmt.Println("Tour of Go: Functions")

	fmt.Printf("add(3, 5): %v\n", add(3, 5))
	fmt.Printf("add2(5, 7): %v\n", add2(5, 7))

	// return multiple values
	a, b := swap("hello", "world")
	fmt.Printf("swap(\"hello\", \"world\"): %q, %q\n", a, b)

	// name rturn value
	c, d := split(17)
	fmt.Printf("split(17): %v, %v\n", c, d)
}