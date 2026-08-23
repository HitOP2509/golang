package main

import "fmt"

func main() {
	// fmt.Println(factorial(4))
	// printNumNToOne(10)
	// printNumOneToN(10)
}

// Find factorial of a number
func factorial(num int) int {
	//Base case
	if num < 0 {
		return 0
	}
	if num == 0 {
		return 1
	}

	//Recursive case
	return num * factorial(num-1)

}

// Write a recursive function that prints numbers from N to 1
func printNumNToOne(num int) {
	if num < 1 {
		return
	}
	fmt.Println(num)
	printNumNToOne(num - 1)
}

// Write a recursive function that prints numbers from 1 to N.
func printNumOneToN(num int) {
	if num < 1 {
		return
	}
	printNumOneToN(num - 1)
	fmt.Println(num)
}
