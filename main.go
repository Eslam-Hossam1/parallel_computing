// Program to illustrate fmt.Print()

package main

// import fmt package
import (
	"fmt"
	"time"
)

func prinNumbers(num int) {
	for i := 0; i <= num; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters(char rune) {
	for i := 'A'; i <= char; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func factorial(num int) int {
	var result int = 1
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func copy_simulation(num int, char rune, done chan bool) {
	for i := num; i >= 1; i-- {
		fmt.Printf("%c", char)
		time.Sleep(time.Second)
	}
	fmt.Println()
	done <- true
}

func main() {

	now := time.Now()
	defer func() {
		fmt.Printf("Excecution time: %d", time.Since(now))
	}()

	signal := make(chan bool)

	var input_num int
	fmt.Printf("Enter Number : ")
	go copy_simulation(20, '/', signal)
	fmt.Scan(&input_num)
	// go copy_simulation(20, '#', signal)
	var result int = factorial(input_num)
	fmt.Printf("result = %d \n", result)

	<-signal

	//go prinNumbers(5)
	//printLetters('E')

}
