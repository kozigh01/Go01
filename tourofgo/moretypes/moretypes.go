package moretypes

import (
	"github.com/mkozigo/go01/shared/print"
)



func MoreTypes() {
	// pointers
	print.PrintHeader("pointers")
	print.PrintSection(pointers1, "pointers")
	
	// structs
	print.PrintHeader("structs")
	print.PrintSection(struct1, "stucts1()")
	print.PrintSection(struct2, "stucts2()")
	print.PrintSection(struct3, "stucts3()")

	// arrays
	print.PrintHeader("arrays")
	print.PrintSection(arrays1, "arrays1()")

	// slices
	print.PrintHeader("slices")
	print.PrintSection(slices1, "first slice")
	print.PrintSection(slices2, "basic slicing")
	print.PrintSection(slices3, "slices3()")
	print.PrintSection(slices4, "slices4()")
	print.PrintSection(slices5, "slices5()")
	print.PrintSection(slices6, "slices6()")
	print.PrintSection(slices7, "slices7()")
	print.PrintSection(slices8, "slices8(): slices of slices")
	print.PrintSection(slices9, "slices9(): appending to slices")
	print.PrintSection(slices10, "slices10(): range of slice")
	print.PrintSection(slices11, "slices11(): slices of slices example")

	// maps
	print.PrintHeader("maps")
	print.PrintSection(maps1, "maps1(): declare, initialize, map literal")
	print.PrintSection(maps2, "maps2(): add, modify, delete")
	print.PrintSection(maps3, "maps3(): maps exercise")

	// functions
	print.PrintHeader("functions")
	print.PrintSection(funcs1, "funcs1(): functions as variables/parameters")
	print.PrintSection(funcs2, "funcs2(): functions closures")
	print.PrintSection(funcs3, "funcs3(): Fibonacci exercise")
}

