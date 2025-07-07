package main

import "fmt"

func increment(val int) {
	val++
	fmt.Println("Incremented value:", val)
}

func incrementPointer(val *int) {
	*val++ // increment the value pointed to by val
	fmt.Println("Incremented value:", *val)
}

func main() {

	val := 10
	// increment(val)                      // passing value, not a pointer
	incrementPointer((&val))
	fmt.Println("Original value:", val) // original value remains unchanged

}
