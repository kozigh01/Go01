package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mkozigo/go01/mascot"
	myprint "github.com/mkozigo/go01/shared/print"
	"github.com/mkozigo/go01/tourofgo"
	"github.com/mkozigo/go01/tourofgo/goroutines"
	"github.com/mkozigo/go01/tourofgo/interfaces"
	"github.com/mkozigo/go01/tourofgo/methods"
	"github.com/mkozigo/go01/tourofgo/moretypes"
	"github.com/mkozigo/go01/tourofgo/functions"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("Hello, world! (time: %s)\n", time.Now())
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	fmt.Println(cmp.Diff("Hello world", "Hello Go"))
	fmt.Println("My favorite number is", rand.Intn(10))

	tourofgo.Basics()

	myprint.PrintDivider()
	tourofgo.FlowOfControl()
	
	myprint.PrintDivider()
	moretypes.MoreTypes()
	
	myprint.PrintDivider()
	functions.Functions()
	
	myprint.PrintDivider()
	moretypes.MoreTypes()
	
	myprint.PrintDivider()
	methods.Methods()

	myprint.PrintDivider()
	interfaces.Interfaces()

	myprint.PrintDivider()
	goroutines.GoRoutines()
}
