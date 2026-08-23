package main

import (
	"fmt"
)

/*********** init function ***********/

// init is a special function in Go that runs automatically before the main function.
// It is mainly used for setup work that should happen before the program starts.
// You do not call init manually. Go calls it automatically.
// A package can have multiple init functions.
// If there are multiple init functions, they run in the order they appear in the file.
// init does not take any parameters and does not return anything.
// NOTE: init runs only once when the package is initialized. So even if you import same package multiple time, init functions will be executed only the first time
// It is commonly used for package-level setup, registering drivers, loading config, or preparing default values.
// Avoid writing too much logic inside init because it can make the program harder to understand and test.

var appName = "My Go App"

func main() {
	fmt.Println("Main function executed. App name is: ", appName)
}

func init() {
	appName = "Init 1"
	fmt.Println("Init function executed")
}

func init() {
	appName = "Init 2"
	fmt.Println("Init 2 function executed")
}

// OUTPUT:
// Init function executed
// Init 2 function executed
// Main function executed. App name is:  Init 2
