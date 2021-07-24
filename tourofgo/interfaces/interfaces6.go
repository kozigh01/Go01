package interfaces

import (
	"fmt"
)

type I6 interface {
	M()
}

func interfaces6() {
	var i I6
	describe6(i)
	i.M()  // calling method on nill interface value throws error
}
func describe6(i I6) {
	fmt.Printf("(%v, %T)\n", i, i)
}