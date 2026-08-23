package main

import (
	"fmt"
	"student-marks-manager/students"
)

func main() {
	var studentid string
	var sub1Marks int
	var sub2Marks int

	fmt.Println("Enter student ID:")
	fmt.Scanln(&studentid)

	fmt.Println("ENTER MARKS FOR SUBJECT ID 1")
	fmt.Scanln(&sub1Marks)
	fmt.Println("ENTER MARKS FOR SUBJECT ID 1")
	fmt.Scanln(&sub2Marks)

	students.UpdateMarks(studentid, "1", sub1Marks)
	fmt.Println("After Subject 1 Marks Update: ", students.Students[studentid])

	students.UpdateMarks(studentid, "2", sub2Marks)
	fmt.Println("After Subject 2 Marks Update: ", students.Students[studentid])

}

func TakeUserInput() (string, int) {
	var marks int
	var studentid string

	fmt.Println("Enter student ID:")
	fmt.Scanln(&studentid)
	fmt.Println("Enter marks:")
	fmt.Scanln(&marks)

	return studentid, marks
}
