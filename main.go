package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const ConferenceTicket = 50
	var remainingTicket uint8 = 50

	bookings := []string{}

	greetUser(conferenceName, ConferenceTicket, remainingTicket)

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValildTicketNumber := userTicket > 0 && userTicket <= remainingTicket

		// isValidCity := city != "Dhaka" && city !=

		if isValildTicketNumber && isValidEmail && isValidName {
			remainingTicket -= userTicket
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTicket, email)
			fmt.Printf("%v tickets reaming for %v. \n", remainingTicket, conferenceName)

			firstNames := []string{}

			for _, name := range bookings {
				var names = strings.Fields(name)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of booking are: %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("we are running out of ticket.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("first name or last name you enter is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("email address does not contains @ sign\n")
			}
			if !isValildTicketNumber {
				fmt.Printf("number of ticket is invalid \n")
			}

			continue
		}

	}
}

func greetUser(conferenceName string, ConferenceTicket uint8, remainingTicket uint8) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", ConferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}
