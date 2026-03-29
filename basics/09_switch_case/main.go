package main

import (
	"fmt"
)

func main() {

	//BY DEFAULT: In Go, you don't have to manually break the switch case,
	// but if you want it to continue executing other cases like other language, use fallthrought

	//Single condition
	var rating int = 5

	switch rating {
	case 1:
		fmt.Println("Very poor")
	case 2:
		fmt.Println("Poor")
	case 3:
		fmt.Println("Average")
	case 4:
		fmt.Println("Good")
	case 5:
		fmt.Println("Very Good")
	default:
		fmt.Println("Invalid Input")
	}

	//Multi condition
	var day = "Sunday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Working Day")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Input")
	}

	//Condition - Do no metion the variable beside switch

	age := 18

	switch {
	case age >= 18:
		fmt.Println("Adult")

	case age >= 16:
		fmt.Println("Teen")

	default:
		fmt.Println("Minor")
	}

	//OUTPUT: with fallthrought
	// Adult
	// Teen
	// Minor

	//OUTPUT: without fallthrought
	// Adult
}
