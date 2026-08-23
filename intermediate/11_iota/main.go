package main

import "fmt"

type Month int

func main() {
	// In Go, iota is a predeclared identifier that acts as an auto-incrementing counter inside const blocks.
	const (
		January Month = 1 + iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)

	fmt.Println(January)  // OUTPUT: 1
	fmt.Println(February) // OUTPUT: 2
	fmt.Println(December) // OUTPUT: 12
}
