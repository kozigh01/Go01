package interfaces

import "math"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}