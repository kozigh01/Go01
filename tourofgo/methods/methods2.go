package methods

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Printf("f.Abs(MyFloat(-math.Sqrt2)): %v\n", f.Abs())
}