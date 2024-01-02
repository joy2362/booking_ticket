package helper

import (
	"fmt"
	"strings"
)

/**
 * validation.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @version	v1.0.1	Tuesday, January 2nd, 2024.
 * @global
 * @param	firstname      	string
 * @param	lastname       	string
 * @param	email          	string
 * @param	userticket     	uint8
 * @param	remainingticket	uint
 * @return	mixed
 */
func Validation(firstName string, lastName string, email string, userTicket uint8, remainingTicket uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValildTicketNumber := userTicket > 0 && userTicket <= remainingTicket
	return isValidName, isValidEmail, isValildTicketNumber
}

/**
 * printValidationError.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @version	v1.0.1	Tuesday, January 2nd, 2024.
 * @global
 * @param	isvalidname         	bool
 * @param	isvalidemail        	bool
 * @param	isvalildticketnumber	bool
 * @return	void
 */
func PrintValidationError(isValidName bool, isValidEmail bool, isValildTicketNumber bool) {
	if !isValidName {
		fmt.Printf("first name or last name you enter is too short\n")
	}
	if !isValidEmail {
		fmt.Printf("email address must be contains @ sign\n")
	}
	if !isValildTicketNumber {
		fmt.Printf("number of ticket is invalid \n")
	}
}
