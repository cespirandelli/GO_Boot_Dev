package main

import "fmt"

// The first three types on the list are the most common for code use and optmization.
// The last one is rarer and will happen only in specific situations.

// int
// uint
// float64
// complex128

// There's a way of converting from integer to float64 (vice versa), as follows.
func main() {
	temperatureInt := 88
	temperatureFloat := float64(temperatureInt) // Casting Float64 on Int type

	fmt.Println("The temperature outside is:", temperatureFloat)
}
