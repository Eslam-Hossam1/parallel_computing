package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 2) //buffered with capacity = 2
	channel <- 500 //sendind
	channel<-600
	fmt.Println(<-channel) //reciving
	fmt.Println(<-channel)

}
