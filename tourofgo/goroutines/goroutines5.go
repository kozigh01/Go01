package goroutines

import (
	"fmt"
)

func fibonacci5(c, quit chan int) {
	x, y := 0, 1

	for {
		fmt.Printf("pushing to channel (before): c <- %v\n", x)
		select {
		case c <- x:
			fmt.Printf("pushing to channel (after): c <- %v\n", x)
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			close(c)
			return
		}
	}
}

func goroutines5() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("pulling from channel: val = %v\n", <-c)
			// fmt.Println(<-c)
		}
		quit <- 0
	}()

	go fibonacci5(c, quit)

	<-quit
}