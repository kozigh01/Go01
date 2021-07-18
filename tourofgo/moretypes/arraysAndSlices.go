package moretypes

import (
	"fmt"
	"strings"
)

func arrays1() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World Foo"
	fmt.Printf("%v: %q, %q\n", a, a[0], a[1])
}

func slices1() {
	primes := [6]int {2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	t := primes[2:5]
	fmt.Printf("primes: %v; primes[1:4]: %v; primes[2:5]: %v\n", primes, s, t)
	fmt.Printf("primes: %v; primes[4:]: %v; primes[:3]: %v\n", primes, primes[4:], primes[:3])
}

func slices2() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxxxx"
	fmt.Printf("%q; %q; %q\n", names, a, b)
}

func slices3() {
	// slice literal - creates underlying array
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct { 
		i int 
		b bool 
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func slices4() {
	a := [6]int{2, 3, 5, 7, 11, 13}
	s := []int{2, 3, 5, 7, 11, 13}


	s = s[:]
	fmt.Println(s, a[:])

	s = s[1:4]
	fmt.Println(s, a[1:4])

	s = s[:2]
	fmt.Println(s, a[:2])

	s = s[1:]
	fmt.Println(s, a[1:])
}

func slices5() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:0]
	printSlice5(s)

	s = s[:4]
	printSlice5(s)

	s = s[2:]
	printSlice5(s)

	s = s[:4]
	printSlice5(s)

	s = append(s, 42)
	printSlice5(s)

	// s = s[:5]  // compiler error, since at this point, the capacity of the slice is only 4
	// printSlice(s)
}
func printSlice5(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

func slices6() {
	var s []int // zero value of slice is nil
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func slices7() {
	a := make([]int, 5)
	printSlice7("a", a)

	b := make([]int, 0, 5)
	printSlice7("b", b)

	b1 := append(b, 1)
	printSlice7("b1", b1)

	b2 := append(b1, 2, 3, 4, 5)
	printSlice7("b", b)
	printSlice7("b1", b1)
	printSlice7("b2", b2)

	b3 := append(b2, 6)
	printSlice7("b3", b3)

	b2[2] = 33
	printSlice7("b", b)
	printSlice7("b1", b1)
	printSlice7("b2", b2)
	printSlice7("b3", b3)

	b2[2] = 3
	b3[2] = 33
	printSlice7("b", b)
	printSlice7("b1", b1)
	printSlice7("b2", b2)
	printSlice7("b3", b3)

	b1[0] = 11
	printSlice7("b", b)
	printSlice7("b1", b1)
	printSlice7("b2", b2)
	printSlice7("b3", b3)

	c := b[:2]
	printSlice7("c", c)
}
func printSlice7(s string, x []int) {
	fmt.Printf("%s: len=%d, cap=%d, %v\n", s, len(x), cap(x), x)
}

func slices8() {
	// slices of slices

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[1][1] = "X"

	for i := 0; i < len(board); i++ 	{
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func slices9() {
	var s []int
	printSlice7("s", s)

	s = append(s, 0)
	printSlice7("s - append 0", s)

	s = append(s, 1)
	printSlice7("s - append 1", s)

	s = append(s, 2, 3, 4)
	printSlice7("s - append 2, 3, 4", s)
}

func slices10() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // 2**1
	}
	for _, v := range pow2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}