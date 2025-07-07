package main

func sum(sums ...int) int {
	total := 0
	for _, v := range sums {
		total += v
	}
	return total

}

func main() {
	// variadic function
	result := sum([]int{1, 2, 3, 4, 5}...) // using the ... operator to unpack the slice
	println("The sum is:", result)
}
