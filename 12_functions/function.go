package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// multiple return values

func addAndSub(a, b int) (int, int) {
	return a + b, a - b
}

// funtion under functions

func math(fn func(a, b int) int) int {
	return fn(10, 5)
}

func main() {
	// basic function

	result := add(5, 10)
	fmt.Println("Result of add function:", result)

	// multiple return values
	sum, sub := addAndSub(10, 5)
	fmt.Println("Sum:", sum, "Subtraction:", sub)

	// function as parameter
	result2 := math(add)
	fmt.Println("Result of math function with add:", result2)

}
