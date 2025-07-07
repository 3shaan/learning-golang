package main

import "fmt"

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// normal for loop
	// for i := 0; i < len(arr); i++ {
	// 	println(arr[i])
	// }

	// now range loop
	// for i, v := range arr {
	// 	fmt.Println("index:", i, "value:", v)
	// }

	// loop in map
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

}
