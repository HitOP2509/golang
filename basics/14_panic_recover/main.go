package main

import "fmt"

/*********** panic ***********/

// panic immediately stops the normal execution of the current function.
// It is similar to throw new Error() in JavaScript.
// When panic is called, the code written after panic will not execute.
// Before the program crashes, Go runs all deferred functions that were encountered before panic.
// panic should not be used for normal errors like invalid input, API failure, or DB query failure.
// It should be used only when the program reaches a state where it cannot safely continue.
// Example: missing required config, failed app startup, or an impossible programmer mistake.

func main() {
	defer fmt.Println("Before panic")
	defer recoverPanic()
	// processOutput(-12)
	processOutput(1)
	defer fmt.Println("After panic. Only executed if processOutput does not panics.") //Even after recovering, this will not executed if processOutput panics
}

func processOutput(input int) {
	if input < 0 {
		panic("Input must be a positive integer")
		// fmt.Println("Unreachable")
	}

	fmt.Println("Valid input: ", input)
}

/*********** recover ***********/

// recover is used to catch a panic and stop the program from crashing.
// It is similar to catch in JavaScript, but it only works inside a deferred function.
// When panic happens, Go starts running deferred functions in Last In First Out order.
// If recover is called inside one of those deferred functions, it catches the panic.
// After recover catches the panic, the program does not continue from the panic line.
// Instead, the surrounding function stops and returns normally.
// recover is mainly useful at app boundaries like HTTP middleware, worker wrappers, or goroutine wrappers.

func recoverPanic() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from panic:", err)
	}
}
