package methods

import "fmt"

func methods4() {
	v := Vertex{3, 4}
	v.Scale(2)  // can call method with a pointer reciever with a value - converts to (&v).Scale
	Scale(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v, p)
}