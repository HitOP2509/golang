package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter your age: ")
	// Use & to pass the address of the variables
	_, err := fmt.Scanln(&age)

	if err != nil {
		fmt.Print("Unable to read input")
		return
	}

	if age >= 18 {
		fmt.Printf("You are an Adult")
	} else if age >= 16 && age < 18 {
		fmt.Printf("You are a Teen")
	} else {
		fmt.Printf("You are a kid")
	}
}
