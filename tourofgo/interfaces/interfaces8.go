package interfaces

import (
	"fmt"
)

func interfaces8() {
	var i interface{} = "hello"
	var ok bool

	s1 := i.(string)
	fmt.Println(s1)
	s1, ok = i.(string)
	fmt.Println(s1, ok)

	i = 42
	// s2 := i.(string) // this throws error because we are not using the second return val
	s2, ok := i.(string)
	fmt.Println(s2, ok)
}