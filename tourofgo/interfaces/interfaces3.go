package interfaces

import (
	"fmt"
)


type I3 interface {
	M()
}


type T3 struct {
	S string
}
 
// this method implements the M() method of the interface I implicitly,
// we do not need to explicitly declare that we are implementing the interface
func (t T3) M() {
	fmt.Println(t.S)
}


func interfaces3() {
	var i I3 = T3{"howdy"}
	i.M()
}