package main

import "fmt"

func main() {
	/*
		Integer Overflow & Underflow in Go

		Go does NOT throw errors for overflow/underflow — values wrap around.

		----------------------------------------
		How wrap-around works (core concept)
		----------------------------------------

		Values behave like a circular range (modulo arithmetic).

		Signed int64 range:
		[-2^63 .......... 0 .......... 2^63-1]

		Overflow (going above max):
		9223372036854775807 + 1 → -9223372036854775808

		Underflow (going below min):
		-9223372036854775808 - 1 → 9223372036854775807

		After wrap:
		- Operations continue normally from the wrapped value

		----------------------------------------

		Unsigned uint64 range:
		[0 .......... 2^64-1]

		Overflow:
		18446744073709551615 + 1 → 0

		Underflow:
		0 - 1 → 18446744073709551615

		After wrap:
		- Values continue incrementing/decrementing normally
	*/

	/*
		---------- OUTPUT OF BELOW CODE: ----------

		Signed max: 9223372036854775807
		After +1 (overflow): -9223372036854775808
		After +2: -9223372036854775807
		Signed min: -9223372036854775808
		After -1 (underflow): 9223372036854775807
		After -2: 9223372036854775806

	*/

	// --- Signed Integer Overflow ---
	var maxInt int64 = 9223372036854775807 // Max int64 (2^63 - 1)
	fmt.Println("Signed max:", maxInt)

	maxInt = maxInt + 1 // Wraps to MIN int64
	fmt.Println("After +1 (overflow):", maxInt)

	maxInt = maxInt + 1
	fmt.Println("After +2:", maxInt)

	// --- Signed Integer Underflow ---
	var minInt int64 = -9223372036854775808 // Min int64 (-2^63)
	fmt.Println("Signed min:", minInt)

	minInt = minInt - 1 // Wraps to MAX int64
	fmt.Println("After -1 (underflow):", minInt)

	minInt = minInt - 1
	fmt.Println("After -2:", minInt)

	/*
		---------- OUTPUT OF BELOW CODE: ----------

		Unsigned max: 18446744073709551615
		After +1 (overflow): 0
		After +2: 1
		Unsigned min: 0
		After -1 (underflow): 18446744073709551615
		After -2: 18446744073709551614

	*/

	// --- Unsigned Integer Overflow ---
	var maxUInt uint64 = 18446744073709551615 // Max uint64 (2^64 - 1)
	fmt.Println("Unsigned max:", maxUInt)

	maxUInt = maxUInt + 1 // Wraps to 0
	fmt.Println("After +1 (overflow):", maxUInt)

	maxUInt = maxUInt + 1
	fmt.Println("After +2:", maxUInt)

	// --- Unsigned Integer Underflow ---
	var minUInt uint64 = 0
	fmt.Println("Unsigned min:", minUInt)

	minUInt = minUInt - 1 // Wraps to MAX uint64
	fmt.Println("After -1 (underflow):", minUInt)

	minUInt = minUInt - 1
	fmt.Println("After -2:", minUInt)
}
