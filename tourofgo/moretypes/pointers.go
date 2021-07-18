package moretypes

import "fmt"

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