package main

import "time"

func main() {
	// simple switch statement

	i := 4

	switch i {
	case 1:
		println("i is 1")
	case 2:
		println("i is 2")

	case 3:
		println("i is 3")

	default:
		println("otherwise")
	}

	// multiple conditions in switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend!")

	default:
		println("It's a weekday.")
	}

	// type switch

	typeSwitch := func(i interface{}) {
		switch i.(type) {
		case int:
			println("i is an int")
		case string:
			println("i is a string")
		case bool:
			println("i is a bool")
		}
	}

	typeSwitch(42)
	typeSwitch("Hello")

}
