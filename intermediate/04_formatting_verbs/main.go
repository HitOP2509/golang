package main

import "fmt"

type User struct {
	Name     string
	Age      int
	Salary   float64
	IsActive bool
	Grade    rune
}

func main() {
	user := User{
		Name:     "Rohit",
		Age:      25,
		Salary:   45678.98765,
		IsActive: true,
		Grade:    'A',
	}

	fmt.Printf("%v\n", user)  // {Rohit 25 45678.98765 true 65}
	fmt.Printf("%+v\n", user) // {Name:Rohit Age:25 Salary:45678.98765 IsActive:true Grade:65}
	fmt.Printf("%#v\n", user) // main.User{Name:"Rohit", Age:25, Salary:45678.98765, IsActive:true, Grade:65}
	fmt.Printf("%T\n", user)  // main.User

	fmt.Printf("%s\n", user.Name) // Rohit
	fmt.Printf("%q\n", user.Name) // "Rohit"

	fmt.Printf("%d\n", user.Age)      // 25
	fmt.Printf("%f\n", user.Salary)   // 45678.987650
	fmt.Printf("%.2f\n", user.Salary) // 45678.99

	fmt.Printf("%t\n", user.IsActive) // true

	fmt.Printf("%c\n", user.Grade) // A
	fmt.Printf("%q\n", user.Grade) // 'A'
	fmt.Printf("%U\n", user.Grade) // U+0041

	fmt.Printf("%p\n", &user) // pointer address, example: 0xc000100000

	fmt.Printf("%d%%\n", user.Age) // 25%
}
