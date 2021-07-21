package methods

import (
	"fmt"
)

func methods1() {	
	v := Vertex{3, 4}
	p := &v
	fmt.Printf("v = %v: v.Abs(): %v\n",v , v.Abs())
	fmt.Printf("p = %v: p.Abs(): %v\n", p, p.Abs())
	fmt.Printf("Abs(v): %v\n", Abs(v))
}