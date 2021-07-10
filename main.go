package main

import (
	"fmt"

	"github.com/kozigh01/go01/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}