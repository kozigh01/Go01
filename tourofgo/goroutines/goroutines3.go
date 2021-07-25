package goroutines

import "fmt"

func goroutines3() {
	ch := make(chan int, 2)
	
	ch <- 1
	ch <- 2
	// ch <- 3  // this will cause a dealock, as it blocks and no go routines are working

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}