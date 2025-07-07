package main

import (
	"fmt"
	"sync"
)

func printTask(i int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that the goroutine is done
	fmt.Println("Task ", i, "is running")
}

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		// go func(i int) {
		// 	fmt.Println("Task ", i, "is running")
		// }(i)

		// wait group
		wg.Add(1)

		go printTask(i, &wg)
	}

	// sleep to allow goroutines to finish
	// time.Sleep(time.Second * 1)

	// wait for all goroutines to finish
	wg.Wait()

}
