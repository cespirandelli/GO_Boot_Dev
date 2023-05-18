// Basic types:

// bool

// string

// Numbers a) (Whole poisitive or negative values): int int8 int16 int32 int64
// Numbers b) (Unsighted integers = Only positive integers): uint uint8 uint16 uint32 uint64 uintptr
//--------------------------------------------------------------------------------------------------//
// --- The number after the "typename_#", represents the number of bits that it can store.
// --- uint8 stores 8 bits of data, whereas uint16 stores 16 bits of data.
// --- The highest number you can store on "uint8" is 255 because it is the hights numbers you can
// --- represent with eight ones and zeroes in binary. Uint16 stores up to the number 65535 and so on.
//--------------------------------------------------------------------------------------------------//
// Numbers c) (Accepts fractional values): float32 float64
// Numbers d) (Used to represent the imaginary numbers): complex64 complex128

// byte (alias for uint8, eight binariy digits)

// rune (alias for int32): Represents a Unicode code point. Something like one character in a string.

// Declaring a variable
// use the "var" keyword. This example the variable is called "number" of type "int".
// -> var number int

// the value of an initialized variable with no assignment will be its "zero value".

// To declare a variable called "pi" to be type "float64"
// -> var pi float64 = 3.14159

// Initialize some variables with their zero values.
package main

import "fmt"

// Initialize here
func main() {
	var smsSendingLimit int // Zero = 0
	var costPerSMS float64  // Zero = 0.000000
	var hasPermission bool  // False
	var username string     // Empty string = ""

	fmt.Println(
		"Sending Limit:", smsSendingLimit,
		"Cost Per SMS:", costPerSMS,
		"Has permission:", hasPermission,
		"Username:", username)
}
