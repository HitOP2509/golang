package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RECOVERED", err)
		}

	}()
	data := []struct {
		COURSE_NAMEE string
	}{
		{COURSE_NAMEE: "Go"},
		{COURSE_NAMEE: "JavaScript"},
	}

	/************ Using template.New() -> We manually need to panic ************/
	// tmpl, err := template.New("Hello").Parse("<h1 data-test-id='Hello'>Welcome to the {{.COURSE_NAME}} course</h1>\n")

	// if err != nil {
	// 	panic(err)
	// }

	// for _, v := range data {
	// 	tmplErr := tmpl.Execute(os.Stdout, v)

	// 	if tmplErr != nil {
	// 		panic(tmplErr)
	// 	}
	// }

	/************ Using template.Must() -> It panics on err ************/

	tmpl := template.Must(template.New("Hello").Parse("<h1 data-test-id='Hello'>Welcome to the {{.COURSE_NAME}} course</h1>\n"))

	for _, v := range data {
		tmplErr := tmpl.Execute(os.Stdout, v)

		if tmplErr != nil {
			panic(tmplErr)
		}
	}

	fmt.Println("TASK COMPLETED!")
}
