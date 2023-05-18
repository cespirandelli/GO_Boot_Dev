package main

import "fmt"

func main() {
	var username string = "Cesar"

	//var password int = 1234567890

	// Concatenate won't happen with missmatched variable types, showing the error below
	// "invalid operation: username + ":" + password (mismatched types string and int)"
	// So it's needed to change the integer to string type first.
	var password string = "1234567890"

	// Concatenate
	fmt.Println("Authorazation: Basic", username+":"+password)
}

// GO programs are easy on memory, they are fairly lightweight.
// Each program includes a small amount of "extra code" that's included in the executable binary.
// This extra code is called "GO Runtime", on of the purposes of it is to cleanup unused memory at runtime.
