package interfaces

import (
	"fmt"
)


type I5 interface {
	M5()
}


type T5 struct {
	S string
}
func (t *T5) M5() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}


func interfaces5() {
	var i I5 

	var t *T5
	i = t
	describe(i)
	i.M5()
	
	i = &T5{"what's up"}
	describe(i)
	i.M5()
}

func describe(i I5) {
	fmt.Printf("(%v, %T)\n", i, i)
}