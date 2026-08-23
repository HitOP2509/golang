package main

import "fmt"

func main() {
	/***************** POINTERS *****************/
	//A pointer is a variable that stores the memory address of another variable

	a := 5
	ptr := &a

	mod := modifyOriginalIntValue(ptr)

	fmt.Println(a)
	mod()
	fmt.Println(a)
	mod()
	fmt.Println(a)
	mod()
	fmt.Println(a)
	mod()
	fmt.Println(a)

}

func modifyOriginalIntValue(i *int) func() (int, int) {
	initialVal := *i
	result := *i
	return func() (int, int) {
		result = *i + initialVal
		*i = result
		return initialVal, result
	}
}
