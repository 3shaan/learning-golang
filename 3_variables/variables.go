package main

import "fmt"

func main() {
	var name string = "John Doe"
	var age int = 30

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// infer type
	var name2 = "Jane Doe"
	fmt.Println("Name2:", name2)

	// short hand
	name3 := "Alice"
	fmt.Println("Name3:", name3)
}
