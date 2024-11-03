package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getMoney(ch chan int) {
	//Accept number of trials
	//محاولات
	var trials int
	fmt.Println("Enter Number of Trails: ")
	if _, err := fmt.Scan(&trials); err != nil {
		fmt.Println("Please Provide Integer number")
		return
	} // Error handling
	rand.Seed(time.Now().UnixNano()) //Pseudu Random
	for i := 1; i <= trials; i++ {
		amount := rand.Intn(500)
		ch <- amount
	}
	//End sending
	close(ch)

	//for knowledge
	// 	infinty loop like while
	// 	for{
	// 		if x > 5{
	// 			break
	// 		}
	// 	}
}

func main() {
	channel := make(chan int)
	go getMoney(channel)
	for msg := range channel { // check whether channel (sending)
		fmt.Printf("You've won: %d$\n", msg)
	}

	// channel <- 500 //sending
	// channel<-600
	// fmt.Println(<-channel) //reciving
	// fmt.Println(<-channel)

}
