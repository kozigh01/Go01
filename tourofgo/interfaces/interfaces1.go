package interfaces

import "fmt"

func interfaces1() {
	//  All methods for a type should be consistent: eiter receive value or receive pointer

	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %v, Abs: %v\n", v, v.Abs())
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	fmt.Printf("Before scaling: %#v, Abs: %v\n", v, v.Abs())
	
	fmt.Println()
	
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())	
}