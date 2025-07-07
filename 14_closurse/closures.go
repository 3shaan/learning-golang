package main

import "fmt"

func counter() func() int {
	sum := 0
	return func() int {
		sum++

		return sum

	}
}

func main() {

	increment := counter()   // create a new counter
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

}
