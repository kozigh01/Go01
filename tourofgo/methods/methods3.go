package methods

import "fmt"

func methods3() {
	v := Vertex{3, 4}
	fmt.Printf("v: %v\n", v)

	// can call Scale with pointer receiver with value type - auto convertes to (&v)
	v.Scale(10)
	fmt.Printf("v.Scale(10): %v\n", v)
	
	v = Vertex{3, 4}
	p := &v
	p.Scale(10)
	fmt.Printf("p = %p, p.Scale(10): %v\n", p, p)
	
	v = Vertex{3, 4}
	Scale(&v, 10)
	fmt.Printf("Scale(v, 10): %v\n", v)
}