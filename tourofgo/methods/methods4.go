package methods

import "fmt"

func methods4() {
	v := Vertex{3, 4}

	// can call method with a pointer reciever with a value - converts to (&v).Scale
	v.Scale(2)
	(&v).Scale(3)
	// can not call a regular function with a pointer parameter with a value
	// Scale(v, 10)  // compile error 'cannot use v (type Vertex) as type *Vertex in argument to Scale'
	Scale(&v, 10)

	p2 := &Vertex{7, 8}
	fmt.Printf("p2 (%T) = %v; p2.Abs() = %v\n", p2, p2, p2.Abs())
	fmt.Printf("*p2 (%T) = %v; (*p2).Abs() = %v\n", *p2, *p2, (*p2).Abs())
	
	x := Abs(*p2)
	fmt.Printf("*p2 (%T) = %v; Abs(*p2) = %v\n", *p2, *p2, x)
	// x := Abs(p2)  // compile error 'cannot use p2 (type *Vertex) as type Vertex in argument to Abs'
}