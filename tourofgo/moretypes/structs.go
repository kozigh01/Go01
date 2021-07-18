package moretypes

import "fmt"

type Vertex struct {
	X int
	Y int
}

func struct1() {
	v1 := Vertex{11, 12}
	fmt.Printf("v1 = %v, X = %v, Y = %v\n", v1, v1.X, v1.Y)
	v1.Y = 33
	fmt.Printf("v1 = %v, X = %v, Y = %v\n", v1, v1.X, v1.Y)
}

func struct2() {
	v1 := Vertex{22, 33}
	var p *Vertex = &v1
	fmt.Printf("p = %p, v1 = %v, X = %v, Y = %v\n", p, *p, p.X, (*p).Y)
}

func struct3() {
	v1 := Vertex{X: 1, Y: 2}
	v2 := Vertex{X: 3} // Y:0 is implicit
	v3 := Vertex{Y: 4} // X:0 is implicit
	v4 := Vertex{} // X:0, Y:0
	p := &Vertex{11, 12}

	fmt.Println(v1, v2, v3, v4, p)
}