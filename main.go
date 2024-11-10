package main
import(
	"fmt"
	"sync"
)

func main() {
	number := [5]int{1,2,3,4,5}
	fmt.Println(number)

}

func printNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num)
}
