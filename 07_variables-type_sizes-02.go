// Assignment:
// Our textio customers want to know how long they have had accounts with us.

// Create a new variable called "accountAgeInt" that will be the truncated,
// integer version of "accountAge".
package main

import "fmt"

func main() {
	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for:", accountAgeInt, "years.")
}
