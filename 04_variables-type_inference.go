// Assignment: "Our current price to send a text message is 2 cents. However, it's likely
// that in the future the price will be a fraction of a penny, so we should use "float64"
// to store this value."

// Create a variable called "penniesPerText" declaration so that it's inferred by the
// compiler to be a "float64".
package main

import "fmt"

func main() {
	var i int         // zero value
	j := i            // j is also an int
	m := 42           // int
	n := 3.14         // float64
	o := 0.867 + 0.5i // complex128
	// penniesPerText := 2 -> This infered to be an integer
	penniesPerText := 2.0 //This infered to be an float64
	fmt.Println(i, j, m, n, o)

	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText) // "The type of penniesPerText is int"
	// the "%T" return the type of the variable, not its value.
	// Instead of using "printLN", use "printF"
}
