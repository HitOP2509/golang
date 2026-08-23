package main

import "fmt"

func counter() (func() int, func()) {
	count := 0

	incrementCount := func() int {
		count += 1
		return count
	}
	resetCount := func() {
		count = 0
	}
	return incrementCount, resetCount
}
func main() {
	incrementCounter, resetCounter := counter()

	incrementCounter()
	incrementCounter()
	incrementCounter()
	incrementCounter()
	fmt.Println(incrementCounter())
	resetCounter()
	fmt.Println(incrementCounter())

}
