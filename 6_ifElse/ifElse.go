package main

func main() {

	age := 18

	if age < 18 {
		println("You are a minor.")
	} else if age >= 18 && age < 65 {
		println("You are an adult.")
	} else {
		println("You are a senior citizen.")
	}

	// varibale in if statement
	if age := 20; age < 18 {
		println("You are a minor.")
	}

	// no ternary operator in Go
}
