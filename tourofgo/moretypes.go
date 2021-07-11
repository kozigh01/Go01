package tourofgo

import "fmt"

type Vertex struct {
	X int
	Y int
}

func MoreTypes() {
	pointers1()
	struct1()
	struct2()
	struct3()
	arrays1()
	slices1()
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
	
}