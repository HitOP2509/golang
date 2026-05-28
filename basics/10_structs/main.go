package main

import "fmt"

type User struct {
	name string
	age  int
}

type TimeEntry struct {
	TaskName string
	Duration int
	Date     string
}

/*
 *	Method with value receiver.
 *	`user User` works like `this` in JS/TS, but Go passes a copy of the User.
 *	Use this when the method only reads data and does not modify the original struct.
 */
func (user User) GetUserName() string {
	return "User name is: " + user.name + "."
}

/*	Method with pointer receiver.
 *	`u *User` works like `this`, but points to the original User.
 *	Use this when the method needs to modify the original struct or avoid copying large structs.
 */
func (u *User) SetName(name string) {
	u.name = name
}

func main() {
	u := User{
		name: "Rohit Singh",
		age:  18,
	}
	task := TimeEntry{
		TaskName: "Coding",
		Duration: 60,
		Date:     "2026-12-12",
	}
	fmt.Println(u.GetUserName())
	u.SetName("Mr. Rohit Singh")
	fmt.Println(u.GetUserName())
	fmt.Println(u.name)

	fmt.Println(task.TaskName)
}
