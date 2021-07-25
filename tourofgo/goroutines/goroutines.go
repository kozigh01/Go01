package goroutines

import (
	"github.com/mkozigo/go01/shared/print"
)

func GoRoutines() {
	print.PrintHeader("go routines")
	// print.PrintSection(goroutines1, "goroutines1")
	// print.PrintSection(goroutines2, "channels")
	// print.PrintSection(goroutines3, "buffered channels")
	// print.PrintSection(goroutines4, "range and close")
	// print.PrintSection(goroutines5, "select")
	// print.PrintSection(goroutines6, "default select")
	print.PrintSection(goroutines7, "sync.Mutex")
}