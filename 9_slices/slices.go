package main

import (
	"fmt"
)

// slices

func main() {

	// uninitialized slice
	var slices []int

	fmt.Println((slices))
	// slices are nill
	fmt.Println(slices == nil)

	// slces with make
	slices2 := make([]int, 5) // [0 0 0 0 0]
	fmt.Println("with make", slices2)
	// capatcity
	fmt.Println("capacity", cap(slices2)) // 5
	// length
	fmt.Println("length", len(slices2)) // 5
	// we can specify capacity also
	slices3 := make([]int, 5, 10) // [0 0 0 0 0]
	fmt.Println("with make and capacity", cap(slices3))

	// slice from array
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slicesFromArr := arr[2:4]
	fmt.Println("slice from array", slicesFromArr) // [3 4]

	// initialize slice with values
	slices4 := []int{1, 2, 3, 4, 5}
	fmt.Println("with values", slices4)

	// append to slice
	slices4 = append(slices4, 6, 7, 8)
	fmt.Println("after append", slices4)

	// change the element
	slices4[3] = 50
	fmt.Println("after change", slices4)

	// append two slices
	slices5 := append(slices3, slices4...)
	fmt.Println("append two slices", slices5)

	// copy
	slices6 := make([]int, len(slices5))
	copy(slices6, slices5)
	fmt.Println("copy slices", slices6)

	// slice of slice
	// slices7 := slices6[5:7]
	slices7 := slices6[:5]                 // [3 4 5]
	fmt.Println("slice of slice", slices7) // [3 4 5]

	// iterate over slice
	for i, v := range slices7 {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// compare two slies
	// nums := []int{1, 2, 3}
	// nums2 := []int{1, 2, 3}
	// fmt.Println("compare", slices.Equal(nums, nums2)) // true

}
