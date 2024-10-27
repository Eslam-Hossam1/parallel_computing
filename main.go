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

func main() {

	now := time.Now()
	defer func() {
		fmt.Printf("Excecution time: %d", time.Since(now))
	}()

	prinNumbers(5)
	printLetters('E')

}
