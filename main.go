package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mkozigo/go01/mascot"
	// "github.com/mkozigo/go01/tourofgo"
	"github.com/mkozigo/go01/tourofgo/moretypes"

	// "github.com/mkozigo/go01/tourofgo/functions"
	"rsc.io/quote"
)

func main() {
	fmt.Printf("Hello, world! (time: %s)\n", time.Now())
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	fmt.Println(cmp.Diff("Hello world", "Hello Go"))
	fmt.Println("My favorite number is", rand.Intn(10))

	// tourofgo.Basics()
	// fmt.Println("\n-------------------------------------\n")
	// tourofgo.FlowOfControl()
	// fmt.Println("\n-------------------------------------\n")
	// tourofgo.MoreTypes()
	// fmt.Println("\n-------------------------------------\n")
	// functions.Functions()
	fmt.Println("\n-------------------------------------\n")
	moretypes.MoreTypes()
}
