package main

import (
	"fmt"
	"strings"
)

// welcomeMessage returns a welcome message for the customer.
func welcomeMessage(customer string) string {
	message := "Welcome to the Tech Palace,"
	return message + " " + strings.ToUpper(customer)
}

// addBorder adds a border to a welcome message.
func addBorder(welcomeMsg string, numStarsPerLine int) string {
	borderStyle := strings.Repeat("*", numStarsPerLine)
	//return boderStyle + "\n" + welcomeMsg + "\n" + boderStyle
	return fmt.Sprintf("%s\n%s\n%s", borderStyle, welcomeMsg, borderStyle)
}

// cleanupMessage cleans up an old marketing message.
func cleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	customer := "Judy"
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	fmt.Println(welcomeMessage(customer))

	fmt.Println(addBorder("Welcome", 10))

	fmt.Println(cleanupMessage(message))
}
