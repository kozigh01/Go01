package methods

import "math"


type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	v.X *= 10
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func(v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// Abs as regular function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Scale as regular function
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
