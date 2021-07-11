package tourofgo

import (
	"fmt"
	"math"
	"math/cmplx"
)

var c, python, java bool
var i, j = 11, 12
var (
	toBe bool = true
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func Basics() {
	c2, python2, java2 := true, false, "no!"
	k := 3

	fmt.Printf("add(4, 5): %d\n", add(4, 5))
	x0, x1 := swap("one", "two")
	fmt.Printf("('one', 'two') swaped: %s, %s\n", x0, x1)
	x, y := swap2("uno", "dos")
	fmt.Printf("('uno', 'dos') swap2'ed: %s, %s\n", x, y)

	
	fmt.Println(c, python, java)
	fmt.Println(i, j)
	fmt.Println(c2, python2, java2, k)

	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s)

	var x2, y2 int = 3, 4
	var f2 float64 = math.Sqrt(float64(x2*x2 + y2*y2))
	var z2 uint = uint(f2)
	fmt.Println(x2, y2, z2)

	v3 := 42
	// v3 := 42.3
	fmt.Printf("Type: %T, Value: %v\n", v3,v3)

	const pi = 3.14159
	// pi = 42  // throws error
	fmt. Println("Pi: ", pi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println(needInt(Big))  // overflow problem
	fmt.Println(needFloat(Big))
	fmt.Printf("Big: type = %T, value = %v\n", float64(Big), float64(Big))
	// fmt.Printf("Big: type = %T, value = %v\n", Big, Big)  // doesn't compile can't use const as value in Printf

}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// using a 'naked' return with named return values
func swap2(a, b string) (x, y string) {
	x = b
	y = a
	return 
}