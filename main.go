package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/joy2362/booking_ticket/helper"
)

var conferenceName string = "Go Conference"

const ConferenceTicket = 50

var remainingTicket uint8 = 50

var bookings []UserData

type UserData struct {
	firstName  string
	lastName   string
	email      string
	userTicket uint8
}

var wg sync.WaitGroup

/**
 * main.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @version	v1.0.1	Tuesday, January 2nd, 2024.
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
			wg.Add(1)
			go sendingEmail(firstName, lastName, email, userTicket)

			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are: %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("we are running out of ticket.")
				break
			}
			defer wg.Wait()
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
 * @version	v1.0.1	Tuesday, January 2nd, 2024.
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
 * @version	v1.0.1	Tuesday, January 2nd, 2024.
 * @global
 * @return	mixed
 */
func getFirstNames() []string {
	firstNames := []string{}

	for _, user := range bookings {
		firstNames = append(firstNames, user.firstName)
	}

	return firstNames
}

/**
 * getUserInput.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Tuesday, January 2nd, 2024.
 * @version	v1.0.1	Tuesday, January 2nd, 2024.
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
 * @version	v1.0.1	Tuesday, January 2nd, 2024.
 * @global
 * @param	firstname 	string
 * @param	lastname  	string
 * @param	email     	string
 * @param	userticket	uint8
 * @return	void
 */
func booking(firstName string, lastName string, email string, userTicket uint8) {
	remainingTicket -= userTicket

	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		userTicket: userTicket,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets reaming for %v. \n", remainingTicket, conferenceName)
}

func sendingEmail(firstName string, lastName string, email string, userTicket uint8) {
	defer wg.Done()
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Printf("##############")
	fmt.Printf("Sending ticket: \n %v \nto email addtess %v \n", email, ticket)
	fmt.Printf("##############")
}
