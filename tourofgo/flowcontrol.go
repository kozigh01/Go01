package tourofgo

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

	x := 333.33
	fmt.Printf("MySqrt(%v): %v, math.Sqrt(%v)\n", x, MySqrt(x), math.Sqrt(x))

	switch1()
	switch2()
	switch3()

	defer1()
	defer2()
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

func switch1() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X!");
	case "linux":
		fmt.Println("Linux!")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch2() {
	fmt.Println("When's Monday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switch3() {
	t := time.Now()
	location, _ := time.LoadLocation("America/Chicago")
	fmt.Printf("Current time: %v\n", t.In(location))
	hour := t.In(location).Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defer1() {
	defer fmt.Println("world")
	fmt.Print("Hello ")
}

func defer2() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v ", i)
	}
	fmt.Println("done")
}