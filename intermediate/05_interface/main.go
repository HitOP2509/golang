package main

import (
	"fmt"
)

/*********** Interface **********/
// We use an interface to define a set of methods that a type must have.

// we use interface as a type and we initialize it's value using the type where we have defined all the methods that the interface has.

// EXAMPLE 1: With struct
type User struct {
	name string
	age  int
}
type UserInterface interface {
	getUserDetails() User
}

func (u User) getUserDetails() User {
	return u
}

// EXAMPLE 2: With int
type Printer interface {
	Print()
}

type Age int

func (a Age) Print() {
	fmt.Println("Age:", a)
}

func (a Age) GetAge() Age {
	return a
}

//Example 3: Real World use case

type Paymenter interface {
	pay(amount float64) (bool, error)
}

func ProcessPayment(amount float64, gatewayName string) (bool, error) {
	fmt.Printf("Payment via %s is processing.\n", gatewayName)
	if amount > 1000000.00 {
		return false, fmt.Errorf("%.2f is more than allowed limit 10,000,00.00", amount)
	}

	return true, nil
}

type Razorpay struct{}

func (r Razorpay) pay(amount float64) (bool, error) {
	res, err := ProcessPayment(amount, "Razorpay")
	return res, err
}

type Cashfree struct{}

func (r Cashfree) pay(amount float64) (bool, error) {
	res, err := ProcessPayment(amount, "Cashfree")
	return res, err
}

type TestingPayment struct{}

func (r TestingPayment) pay(amount float64) (bool, error) {
	res, err := ProcessPayment(amount, "TestingPayment")
	return res, err
}

func Checkout(p Paymenter, amount float64) (bool, error) {
	if p == nil {
		panic("Paymenter is invalid.")
	}
	result, err := p.pay(amount)
	return result, err
}

func main() {
	var age Printer = Age(25)
	var u1 UserInterface = User{
		name: "Rohit Singh",
		age:  28,
	}

	fmt.Printf("%+s\n", u1.getUserDetails().name) // Rohit Singh
	age.Print()                                   // Age: 25
	// fmt.Printf("%+v\n", age.GetAge()) -> This won't work as age is Printer type and not Age type

	result, err := Checkout(TestingPayment{}, 1000000)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Payment success:", result)

}
