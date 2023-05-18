// Assignment:
// "Within the same function, declare a float called "averageOpenRate" and string called "displayMessage"
// on the same line."

// "Initialize them to values of ".23" and "is the average open rate of your messages" respectively
// before they are printed."

package main

import "fmt"

func main() {
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
}
