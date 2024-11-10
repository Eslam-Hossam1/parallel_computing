package main

import (
	"fmt"
	"sync"
)

func main() {
	number := [5]int{1, 2, 3, 4, 5} // [1,2,3,4,5]
	fmt.Println(number)
	slice := number[1:4] // len = 3, capacity = from 1 (first number in slice brackets  or initial whatever) to the end of original array = 4

	fmt.Println(" slice len :", len(slice))
	fmt.Println(" slice  cap:", cap(slice))

	slice[0] = -10
	fmt.Println("new slice : ", slice)
	fmt.Println("new arr : ", number)
//---------------------------------------------------------------------------------------------------
	type Student struct{
		courses [5]string
		grades [5]float64
	}

	qassas := Student{
		courses :[5]string{"DB", "IT","AI","ML","CN"},
		grades: [5]float64{85.0, 90.0,92.0,95.5,100.0},
	}

	fmt.Println("course: ", qassas.courses)
}

func printNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num)
}
