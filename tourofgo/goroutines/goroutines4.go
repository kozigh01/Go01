package goroutines

import (
	"fmt"
)

func fibonacci4(n int, c chan int) {
	x, y := 0, 1

	for i:= 0; i < n; i++ {
		fmt.Printf("pushing to channel (before): c <- %v\n", i)
		c <- x
		fmt.Printf("pushing to channel (after): c <- %v\n", i)
		x, y = y, x+y
	}

	fmt.Println("closing the channel")
	close(c)
}

func goroutines4() {
	cap := 10
	// c := make(chan int)
	// c := make(chan int,  1)
	c := make(chan int, 10)

	go fibonacci4(cap, c)
	for i := range c {
		fmt.Printf("pulling from channel: val = %v\n", i)
	}
}