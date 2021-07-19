package methods

import (
	"fmt"
)

func methods1() {	
	v := Vertex{3, 4}
	p := &v
	fmt.Printf("v.Abs(): %v\n", v.Abs())
	fmt.Println(v)
	fmt.Printf("p.Abs(): %v\n", p.Abs())
	fmt.Println(p)
	fmt.Printf("Abs(v): %v\n", Abs(v))
}