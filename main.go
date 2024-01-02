package main

import (
	"fmt"
	"strings"

	"github.com/joy2362/booking_ticket/helper"
)

var conferenceName string = "Go Conference"

const ConferenceTicket = 50

var remainingTicket uint8 = 50

var bookings []string

/**
 * main.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @global
 * @return	void
 */
func main() {
	greetUser()

	for {
		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValildTicketNumber := helper.Validation(firstName, lastName, email, userTicket, remainingTicket)

		if isValildTicketNumber && isValidEmail && isValidName {
			booking(firstName, lastName, email, userTicket)
			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are: %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("we are running out of ticket.")
				break
			}
		} else {
			helper.PrintValidationError(isValidName, isValidEmail, isValildTicketNumber)
			continue
		}
	}
}

/**
 * greetUser.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @global
 * @return	void
 */
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", ConferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}

/**
 * getFirstNames.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @global
 * @return	mixed
 */
func getFirstNames() []string {
	firstNames := []string{}

	for _, name := range bookings {
		var names = strings.Fields(name)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

/**
 * getUserInput.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @global
 * @param	mixed	string
 * @param	mixed	string
 * @param	mixed	string
 * @param	mixed	uint8
 * @return	mixed
 */
func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint8

	fmt.Printf("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email:")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

/**
 * booking.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @global
 * @param	firstname 	string
 * @param	lastname  	string
 * @param	email     	string
 * @param	userticket	uint8
 * @return	void
 */
func booking(firstName string, lastName string, email string, userTicket uint8) {
	remainingTicket -= userTicket

	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets reaming for %v. \n", remainingTicket, conferenceName)
}
