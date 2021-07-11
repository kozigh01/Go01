package tourofgo

import "fmt"

type Vertex struct {
	X int
	Y int
}

func MoreTypes() {
	pointers1()
	
	println()
	struct1()

	println()
	struct2()

	println()
	struct3()

	println()
	arrays1()

	println()
	slices1()

	println()
	slices2()

	println()
	slices3()

	println()
	slices4()

	println()
	slices5()

	println()
	slices6()

	println()
	slices7()
}

func pointers1() {
	i, j := 42, 2701

	p := &i
	fmt.Printf("p = %v, *p = %v\n", p, *p)
	*p = 21
	fmt.Printf("p = %v, *p = %v\n", p, *p)
	
	p = &j
	*p = *p / 37
	fmt.Printf("p = %v, *p = %v\n", p, *p)
}

func struct1() {
	v1 := Vertex{11, 12}
	fmt.Printf("v1 = %v, X = %v, Y = %v\n", v1, v1.X, v1.Y)
}

func struct2() {
	v1 := Vertex{22, 33}
	p := &v1
	fmt.Printf("p = %p, v1 = %v, X = %v, Y = %v\n", p, *p, p.X, (*p).Y)
}

func struct3() {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 3} // Y:0 is implicit
	v3 := Vertex{Y: 4} // X:0 is implicit
	v4 := Vertex{} // X:0, Y:0
	p := &Vertex{11, 12}

	fmt.Println(v1, v2, v3, v4, p)
}

func arrays1() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World Foo"
	fmt.Printf("%q, %q\n", a[0], a[1])
	fmt.Println(a)
}

func slices1() {
	primes := [6]int {2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	t := primes[2:5]
	fmt.Println(s, t)
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