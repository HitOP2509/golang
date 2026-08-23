package main

import "fmt"

func main() {
	// We can use Double quotes ("") or Backticks (``) to create string.
	// We use "" when we want to use escape characters. And when we use ``, escape characters will be stored as it is.
	// NOTE: We use Single quotes to create a rune literal.  A rune is an alias for int32 and represents a Unicode code point.

	normalStr := "Hello,\nGo!"
	literalStr := `Hello,\nGo!`
	runeChar := 'A'

	fmt.Println(normalStr, "\n", literalStr)
	fmt.Println(runeChar)
	//OUTPUT:
	// Hello,
	// Go!
	// Hello,\nGo!
	// 65
}
