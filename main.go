package main

import(
	"fmt"
)

func main(){
	channel := make(chan int)
	channel<-500 //sending
	fmt.Println(<-channel) //reciving
	
}