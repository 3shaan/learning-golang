package main

import "fmt"

// func printSlice(nums []int) {
// 	for _, v := range nums {
// 		fmt.Println(v)
// 	}
// }

// func printStringSlice(names []string) {
// 	for _, v := range names {
// 		fmt.Println(v)
// 	}
// }

// // generetic to print slice
// func printGenericSlice[T comparable](s []T) {
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

func slicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if s[i] == v {
			return i
		}
	}
	return -1
}

// lets make a linked list

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next  *element[T]
	value T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: v}
		lst.tail = lst.tail.next
	}
}

// iteretor for the linked list

func (lst *List[T]) allElemets() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func main() {

	// nums := []int{1, 2, 3, 4, 5}
	// printSlice(nums)

	// names := []string{"Alice", "Bob", "Charlie", "Diana"}
	// printStringSlice(names)
	// printGenericSlice(nums)

	// get the index of a value in a slice
	index := slicesIndex([]int{1, 2, 3, 4, 5}, 3)
	if index == -1 {
		println("Value not found in slice")
	} else {
		println("Index of 9 in slice:", index)
	}

	lst := List[int]{}
	lst.Push(2)
	lst.Push(4)
	lst.Push(6)
	fmt.Println("All elements in the list:", lst.allElemets())
}
