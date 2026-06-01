package main

import (
	"errors"
	"fmt"
)

var m map[string]int = map[string]int{"1": 1, "2": 2}

// 1. function returning multiple values
func findResult(id string) (int, bool) {
	value, hasValue := m[id]

	return value, hasValue
}

// 2. function returning error
func isInt(value any) (int, error) {
	val, ok := value.(int)
	if !ok {
		return val, errors.New("Not an int")
	}
	return val, nil
}

// 3. Variadic function - function that can recieve n number of args
func sum(nums ...any) (int, error) {
	result := 0
	var err error = nil

	for _, num := range nums {
		val, e := isInt(num)
		if e != nil {
			err = e
			break
		}
		result += val
	}

	return result, err
}

func main() {
	//Func 1
	a, b := findResult("100")
	fmt.Println(a, b)

	//Func 3
	nums := []any{1, 2, 3, 4, 5, 6}
	res, err := sum(nums...)
	if err != nil {
		fmt.Println("Some values are not integer.")
		return
	}
	fmt.Println(res)
}
