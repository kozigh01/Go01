package tourofgo

import (
	"fmt"
	"math"
)

func FlowOfControl() {
	forLoop1()
	forLoop2()
	// // loop forever
	// for {

	// }

	// if examples
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))

	x := 33.33
	fmt.Printf("MySqrt(%v): %v, math.Sqrt(%v)\n", x, MySqrt(x), math.Sqrt(x))
}

func forLoop1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("%v ", sum)
	}
	fmt.Printf("\nFor loop 1: sum = %v\n", sum)
}

func forLoop2() {
	sum := 1

	// for is golang's while loop
	for sum < 1000 {
		sum += sum
		fmt.Printf("%v ", sum)
	}
	fmt.Printf("\nFor loop 2: sum = %v\n", sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim  // note can't use v here, it is out of scope
}

func MySqrt(x float64) float64 {
	z := 1.0

	fmt.Printf("Sqrt(%v):\n", x)
	for i := 0; i < 10; i++ {
		fmt.Printf("   | i = %v, z = %v\n", i, z)
		z -= (z*z - x) / (2*z)
	}

	fmt.Printf("   final z = %v\n", z)
	return z
}