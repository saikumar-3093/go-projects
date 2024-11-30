package helper

import "strings"

func ValidateUserInput(userFirstName string, userLastName string, userEmail string, ticketsBooked uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := ticketsBooked > 0 && ticketsBooked <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
