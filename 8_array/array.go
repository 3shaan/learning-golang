package main

import "fmt"

func main() {
	// int array

	// 0 numbered values
	var nums [5]int
	nums[0] = 1
	fmt.Println(nums) // [1 0 0 0 0]

	// boolian array
	var bools [3]bool
	bools[0] = true
	fmt.Println(bools) // [true false false]

	// string array
	var names [3]string
	names[0] = "Alice"
	fmt.Println(names) // [Alice  ]

	// array with values
	var nums2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums2) // [1 2 3 4 5]

	// 2D array
	matrix := [2][3]int{{1, 2}, {4, 5, 6}}
	fmt.Println(matrix) // [[1 2 0] [4 5 6]]

	// when use array
	// - when you know the size of the array in advance
	// - when you want to store a fixed number of elements of the same type
	// - optiomized for performance and memory usage

}
