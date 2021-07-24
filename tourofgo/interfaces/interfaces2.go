package interfaces

import (
	"fmt"
	"math"
)


type Abser interface {
	Abs() float64
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func interfaces2() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Printf("f (%T) = %v: a (%T) = %v\n", f, f, a, a)
	fmt.Printf("a = f; a.Abs() = %v\n", a.Abs())
	
	a = &v // a *Vertex implements Abser
	fmt.Printf("v (%T) = %v: a (%T) = %v\n", v, v, a, a)
	fmt.Printf("a = &v; a.Abs() = %v\n", a.Abs())

	// a = v // compile error: v is a Vertex, not a *Vertext: so, doesn't meet requirements for Abs

	fmt.Println(a.Abs())
}