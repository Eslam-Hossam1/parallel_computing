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

func main() {

	now := time.Now()
	defer func() {
		fmt.Printf("Excecution time: %d", time.Since(now))
	}()
	var input_num int
	fmt.Printf("Enter Number : ")
	fmt.Scan(&input_num)
	var result int = factorial(input_num)
	fmt.Printf("result = %d \n", result)
	//go prinNumbers(5)
	//printLetters('E')

}
