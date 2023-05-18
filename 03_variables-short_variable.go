// There is an easier way of declaring variables.
// Lets use the SHORT ASSIGNMENT OPERATOR
// :=

// Assignment:
// "Declare a variable named "congrats" with the value "happy birthday!" using a short variable declaration."
package main

import "fmt"

func main() {
	var empty string    // Empty string without hiding full varialble declaration.
	empty2 := ""        // This is the same as "empty", but written in a condensed way.
	numCars := 10       // inferred to be an integer
	temperature := 0.0  // Temperature is inferred to be a floating number.
	var isFunny = true  // isFunny is inferred to be a boolean
	isitThough := false // inferred to be a boolean shortly

	congrats := "Happy Birthday!"

	fmt.Println(empty, empty2, numCars, temperature, isFunny, isitThough)
	fmt.Println("Assignment:", congrats)
}

// Outside of a function (in the global/package scope), every statement begins with a keyword
// "var", "func"
// In this case it's not possible to use ":=" (short assignment operator)
