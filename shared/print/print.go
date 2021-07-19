package print

import "fmt"

func PrintHeader(header string) {
	fmt.Printf("\n------  %v  -----\n", header)
}

func PrintSection(f func(), description string) {
	fmt.Printf("\n-- %s\n", description)
	f()
}