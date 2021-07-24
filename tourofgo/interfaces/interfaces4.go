package interfaces

import (
	"fmt"
	"math"
)


type I4 interface {
	M()
}

type T4 struct {
	S string
}
func (t *T4) M() {
	fmt.Println(t.S)
}


type F4 float64
func (f F4) M() {
	fmt.Println(f)
}


func interfaces4() {
	var i I4

	i = &T4{"hello"}
	describe4(i)
	i.M()

	i = F4(math.Pi)
	describe4(i)
	i.M()
}

func describe4(i I4) {
	fmt.Printf("(%v, %T)\n", i, i)
}