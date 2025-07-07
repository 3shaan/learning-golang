package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps with make
	m := make(map[string]int)

	m["one"] = 1

	m["two"] = 2

	//clear maps
	// clear(m)

	// delete a key
	// delete(m, "one")

	fmt.Println("map with make:", m)
	fmt.Println("map value without initialze key:", m["three"]) // zero value for int is 0, empty string for string, false for bool

	// make with values
	m2 := map[string]int{
		"age":    30,
		"height": 180,
	}
	fmt.Println("map with values:", m2)

	// check value exists

	v, ok := m2["age"]
	fmt.Println(v, ok)

	if ok {
		fmt.Println("key 'age' exists in m2")
	} else {
		fmt.Println("key 'age' does not exist in m2")
	}

	// equal checks
	fmt.Println(maps.Equal(m, m2)) // false, because keys are different
}
