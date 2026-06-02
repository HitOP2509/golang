package main

import (
	"fmt"
	"os"
)

/*********** exit ***********/

// os.Exit immediately terminates the program with the provided exit code.
// It stops the program instantly and does not wait for the surrounding function to complete.
// NOTE: deferred functions are NOT executed when os.Exit is called.
// Exit code ranges from 0-125. 0 usually means the program completed successfully.
// Non-zero exit codes usually mean the program failed because of some error.
// os.Exit should mainly be used when you intentionally want to stop the whole program immediately.
// Example: missing required config, failed startup validation, or CLI command failure.

var CONFIG map[string]string = map[string]string{
	// "env": "prod",
}

func main() {
	defer func() {
		fmt.Println("Deferred statement") // If get config !ok, this will not be executed
	}()

	conf, ok := getConfig("env")
	if !ok {
		os.Exit(1)
	}
	fmt.Println(conf)
}

func getConfig(key string) (string, bool) {
	val, exists := CONFIG[key]
	return val, exists
}

// OUTPUT
// exit status 1
