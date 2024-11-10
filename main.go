package main
import(
	"fmt"
	"sync"
)

func main() {
	number := [5]int{1,2,3,4,5} // [1,2,3,4,5]
	fmt.Println(number)

	slice := number[1:4] // len = 3, capacity = from 1 (first number  or initial whatever) to the end of original array = 4
	
	fmt.Println(slice)

}

func printNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num)
}
