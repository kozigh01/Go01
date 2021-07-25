package interfaces

import (
	"fmt"
	"math"
)

// type ErrNegSqrt struct {
// 	Value float64
// }
type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegSqrt(x)
	}

	return math.Sqrt(x), nil
}

func interfaces13() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}