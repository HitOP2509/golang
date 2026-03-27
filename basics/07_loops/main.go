package main

import "fmt"

func main() {
	//For loop - Normal
	const MAX = 10
	sum := 0
	for i := 0; i < MAX; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	//For loop as While loop & The init and post statements are optional.
	sum2 := 0
	for sum2 < 10 {
		sum2++
	}
	fmt.Println("Sum2:", sum2)

	//For with range - we can use range to iterate of Arays or Slices
	for i, j := range []int{1, 2, 3} {
		fmt.Printf("Value is %d for Index %d\n", j, i)
	}

	//For with range - we can use range to iterate numbers
	// 	This syntax is new in Go 1.22+.

	// Before that, range only worked on: slices,arrays,maps,strings
	for i := range 3 {
		//i starts from 0 (treated as Index)
		fmt.Println(i)
	}
}
