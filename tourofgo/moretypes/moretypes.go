package moretypes

import "fmt"


func printHeader(header string) {
	fmt.Printf("\n------  %v  -----\n", header)
}

func MoreTypes() {
	// pointers
	// printHeader("pointers")
	// pointers1()
	
	// structs
	// printHeader("structs")
	// struct1()

	// println()
	// struct2()

	// println()
	// struct3()

	// arrays
	// printHeader("arrays")
	// arrays1()

	// slices
	printHeader("slices")
	// printSection(slices1, "first slice")
	// printSection(slices2, "basic slicing")
	// printSection(slices3, "slices3()")
	// printSection(slices4, "slices4()")
	// printSection(slices5, "slices5()")
	// printSection(slices6, "slices6()")
	// printSection(slices7, "slices7()")
	// printSection(slices8, "slices8(): slices of slices")
	// printSection(slices9, "slices9(): appending to slices")
	printSection(slices10, "slices10(): range of slice")
}

func printSection(f func(), description string) {
	fmt.Printf("\n-- %s\n", description)
	f()
}
