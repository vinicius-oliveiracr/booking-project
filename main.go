package main

import (
	"fmt"
	"strings"
	"time"
)

var conferenceName = "Go conference"

const conferencePrice = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, validTicketNumber := validateInputs(firstName, lastName, userEmail, userTickets)

		if isValidName && isValidEmail && validTicketNumber {

			bookTicket(userTickets, firstName, lastName, userEmail)

			firstNames := getFirstNames()
			fmt.Printf("The first names of our bookings are: %v. \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("First or last name is invalid. \n")
			}
			if !isValidEmail {
				fmt.Printf("Email provided does not contain '@' sign. \n")
			}

			if !validTicketNumber {
				fmt.Printf("Number of tickets is invalid. \n")
			}
		}
	}

}

func greetUser() {
	fmt.Printf("Get your tickets to %v now!\n", conferenceName)
	fmt.Printf("At the moment, we have %v tickets available.\n", remainingTickets)
	fmt.Printf("You can buy a ticket for %v dollars.\n", conferencePrice)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func validateInputs(firstName string, lastName string, userEmail string, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) > 2 && len(lastName) > 2

	isValidEmail := strings.Contains(userEmail, "@") && strings.Contains(userEmail, ".")

	validTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, validTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var userEmail string
	var userTickets uint
	var lastName string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Type here your best email:")
	fmt.Scan(&userEmail)

	fmt.Println("Type how many tickets do you want to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     userEmail,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", userData.firstName, userData.lastName, userData.tickets, userData.email)
	fmt.Printf("%v tickets still available for %v. \n", remainingTickets, conferenceName)

	go sendTicket(userTickets, firstName, lastName, userEmail)
}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, userEmail)
	fmt.Println("##################")
}
