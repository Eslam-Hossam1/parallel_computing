package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 1)
	channel <- 500         //sending
	fmt.Println(<-channel) //reciving

}
