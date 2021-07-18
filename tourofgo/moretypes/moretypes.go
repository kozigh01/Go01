package moretypes

import "fmt"


func printHeader(header string) {
	fmt.Printf("\n------  %v  -----\n", header)
}

func MoreTypes() {
	// pointers
	// printHeader("pointers")
	// printSection(pointers1, "pointers")
	
	// structs
	// printHeader("structs")
	// printSection(struct1, "stucts1()")
	// printSection(struct2, "stucts2()")
	// printSection(struct3, "stucts3()")

	// arrays
	// printHeader("arrays")
	// printSection(arrays1, "arrays1()")

	// slices
	// printHeader("slices")
	// printSection(slices1, "first slice")
	// printSection(slices2, "basic slicing")
	// printSection(slices3, "slices3()")
	// printSection(slices4, "slices4()")
	// printSection(slices5, "slices5()")
	// printSection(slices6, "slices6()")
	// printSection(slices7, "slices7()")
	// printSection(slices8, "slices8(): slices of slices")
	// printSection(slices9, "slices9(): appending to slices")
	// printSection(slices10, "slices10(): range of slice")
	// printSection(slices11, "slices11(): slices of slices example")

	// maps
	printHeader("maps")
	// printSection(maps1, "maps1(): declare, initialize, map literal")
	printSection(maps2, "maps2(): add, modify, delete")
}

func printSection(f func(), description string) {
	fmt.Printf("\n-- %s\n", description)
	f()
}
