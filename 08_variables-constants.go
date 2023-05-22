// Keyword: const

// Constants are static value variables.
// It's not possible to use short declaration syntax ":=" with constants.
// Constants can be character, string, boolean or numeric values.
// Constants cannot be complex types as slices, maps and structs.

// Assignment:
// "Create 2 separate constants:
// "premiumPlanName" and "basicPlanName"."

package main

import "fmt"

func main() {
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("Plan:", premiumPlanName)
	fmt.Println("Plan:", basicPlanName)

}
