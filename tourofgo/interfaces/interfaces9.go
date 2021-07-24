package interfaces

import (
	"fmt"
)

type T9 struct {
	S string
}

func interfaces9() {
	do(21)
	do("hello")
	do(true)
	do(T9{"hello"})
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v = %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("%v: %v && false = %v\n", v, v, v && false)
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}