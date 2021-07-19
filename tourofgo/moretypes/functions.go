package moretypes

import (
	"fmt"
	"math"
)

func funcs1() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("Running hypot(3, 4): %v\n", hypot(3,4))
	
	fmt.Printf("Running compute(hypot): %v\n", compute(hypot))
	fmt.Printf("Running compute(math.pow): %v\n", compute(math.Pow))

	
}
func compute(f func(float64, float64) float64) float64 {
	return f(3, 4)
}

func funcs2() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Printf("pos: %v, neg: %v\n", pos(i), neg(-2*i))
	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func funcs3() {
	f := fibonacci()
	
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func fibonacci() func() int {
	fibiteration := 0
	fibcurrent := 0
	fiblast1 := 1
	fiblast2 := 1

	return (func() int {
		switch fibiteration {
			case 0: 
				fibiteration += 1
				return 0
			case 1, 2:
				fibiteration += 1
				return 1
			default:
				fibcurrent = fiblast1 + fiblast2
				fiblast2 = fiblast1
				fiblast1 = fibcurrent
				return fibcurrent
			}
	})
	
}