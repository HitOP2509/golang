package main

import (
	"errors"
	"fmt"
	"math"
)

func isSqrt(n int) (int, error) {
	var val int
	if n < 0 {
		return val, errors.New("Negative value")
	}
	if n == 0 {
		return val, errors.New("Can't calcualte sqrt of 0")
	}
	val = int(math.Pow(float64(n), float64(n)))

	return val, nil
}

func processData[T any](data T) error {
	allowedTypes := map[string]bool{
		"string": true,
		"int":    true,
	}

	dataType := fmt.Sprintf("%T", data)

	_, ok := allowedTypes[dataType]

	if !ok {
		return fmt.Errorf("ERROR: Invalid data type. Only string and int types are supported")
	}

	return nil
}
func main() {
	val1, err1 := isSqrt(1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println(val1)

	if err := processData(true); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data processed.")
}
