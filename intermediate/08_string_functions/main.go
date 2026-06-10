package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "Rohit"
	strHin := "हेलो"

	// NOTE: In Go, str[0] gives a byte, not the actual character.

	fmt.Println(str[0])        //OUTPUT will be in ASCII: 82
	fmt.Printf("%c\n", str[0]) //OUTPUT: R
	fmt.Println(str[0:1])      //OUTPUT: R (When accessing range of a string)

	//So for Hindi chars, it will break because Hindi charcaters use multiple bytes in UTF-8.
	//For that use this
	runes := []rune(strHin)
	fmt.Printf("%c\n", runes[0]) //OUTPUT: ह

	//String conversion
	num := 10
	strNew := "10"
	fmt.Println(strconv.Itoa(num)) // int to string  -> returns string

	val, _ := strconv.Atoi(strNew) // string to int  -> returns int, err
	fmt.Println(val == num)        // OUTPUT: true

	//String split
	fmt.Println(strings.Split(str, "o")) // OUTPUT: [R hit]

	//String joining
	arr := []string{"Apple", "Banana"}
	fmt.Println(strings.Join(arr, " | ")) //OUTPUT: Apple | Banana

	// Regex
	const OTP string = "123456"
	matched, err := regexp.MatchString(`^\d{6}$`, OTP)
	if err != nil {
		fmt.Println("Regex error:", err)
		return
	}

	fmt.Println(matched) // true

	// String Builder -> Use this when memory efficiency is required
	// Instead of creating multiple temp string, it creates a buffer and keeps on epxanding the string in the same buffer instead of new allocation. However, it may also create new buffer to grow the size. To fix the issue, we need to pre-assign the capacity using .Grow

	arrStr := []string{"Hello", "World!", "How", "Are", "You?"}

	var builder strings.Builder

	builder.Grow(len(arrStr)) //Pre-allocating buffer size

	for _, v := range arrStr {
		builder.WriteString(v)
		builder.WriteString(" ")
	}

	//Converting builder to string
	result := builder.String()

	fmt.Println(result)

	//Reset builder -> It clears the content, drops the buffers so The garbage collector can clean it later if there are no other references.
	fmt.Println(builder.Cap())
	builder.Reset()
	fmt.Println(builder.Cap())
}
