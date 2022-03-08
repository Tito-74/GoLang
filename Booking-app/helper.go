package main

import "strings"

// , remainingTickets uint
func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidemail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidemail, isValidTicketNumber

}
