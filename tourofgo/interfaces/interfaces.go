package interfaces

import (
	"github.com/mkozigo/go01/shared/print"
)

func Interfaces() {
	print.PrintHeader("interfaces")
	// print.PrintSection(interfaces1, "interfaces1()")
	// print.PrintSection(interfaces2, "basic interface implementation")
	// print.PrintSection(interfaces3, "interfaces are implemented implicitly")
	// print.PrintSection(interfaces4, "interfaces are like tuple (value, type)")
	// print.PrintSection(interfaces5, "interfaces and nil values")
	// print.PrintSection(interfaces6, "nil interface values")
	// print.PrintSection(interfaces7, "the empty interface")
	// print.PrintSection(interfaces8, "type assertions")
	print.PrintSection(interfaces9, "type switches")
}