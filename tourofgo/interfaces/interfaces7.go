package interfaces

import (
	"fmt"
	"math"
)

func interfaces7() {
	var i interface{} // define i as empty interface, values of any type can be assigned
	describe7(i)

	i = 42
	describe7(i)

	i = math.Pi
	describe7(i)

	i = false
	describe7(i)

	i = "hello world"
	describe7(i)
}
func describe7(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}