package main

import "fmt"

/*********** defer ***********/

// defer keyword deferrs the execution of the code untill the sorrounding functions execution is completed.
// It works as a stack, Last In First Out. If we have multiple defer statement inside a function, it will be executed in LIFO
// NOTE: argument values are identified as soon as defer statement is encountered. So even if it gets executed at the end, it remembers the value when defer statement was initially encountered

func main() {
	count := 1
	defer logAuditTrail(fmt.Sprintf("Trail %v", count))
	fmt.Printf("Count %v\n", count)
	defer logAuditTrail(fmt.Sprintf("Trail %v", count))
	count += 1
	fmt.Printf("Count %v\n", count)
}

//OUTPUT:
// Count 1
// Count 2
// Audit trail logging demo. Trail 1
// Audit trail logging demo. Trail 1 => Despite getting logged after the count was updated, the value still holds the reference of original value of count when defer statement was encountered.

func logAuditTrail(message string) {
	fmt.Println("Audit trail logging demo.", message)
}
