package common

import (
	"strings"
)

const (
	ConferenceTickets uint = 50
)

func ValidateUserInput(firstName string, lastName string, email string, tickets uint, conferenceRemainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := tickets > 0 && tickets <= conferenceRemainingTickets

	return isValidName, isValidEmail, isValidTickets
}