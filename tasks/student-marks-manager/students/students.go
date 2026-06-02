package students

import "fmt"

type Result struct {
	SubjectId string
	Marks     int
}

type Student struct {
	Id      string
	Name    string
	Class   string
	Results []Result
}

var Students = map[string]*Student{
	"1": {
		Id:      "1",
		Name:    "Rohit Singh",
		Class:   "8",
		Results: []Result{},
	},
	"2": {
		Id:      "2",
		Name:    "Mohit Singh",
		Class:   "8",
		Results: []Result{},
	},
}

func UpdateMarks(studentId string, subjectId string, marks int) {
	student := FindStudent(studentId)

	if student == nil {
		fmt.Printf("Student id %s doesn't exist", studentId)
		return
	}

	result := FindResult(student, subjectId)

	if result == nil {
		student.Results = append(student.Results, Result{
			SubjectId: subjectId,
			Marks:     marks,
		})
		return
	}
	result.Marks = marks
}

func FindStudent(studentId string) *Student {
	student, exists := Students[studentId]

	if !exists {
		return nil
	}

	return student
}

func FindResult(student *Student, subjectId string) *Result {
	var resultPointer *Result = nil

	for index, value := range student.Results {
		if value.SubjectId == subjectId {
			resultPointer = &student.Results[index]
		}
	}

	return resultPointer
}
