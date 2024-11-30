package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName     string
	lastName      string
	email         string
	ticketsBooked uint
}

var wg = sync.WaitGroup{}

func main() {

	var ticketsBooked uint
	var userFirstName string
	var userLastName string
	var userEmail string

	greetUser()

	userFirstName, userLastName, userEmail, ticketsBooked = userInput(userFirstName, userLastName, userEmail, ticketsBooked)

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userFirstName, userLastName, userEmail, ticketsBooked, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(ticketsBooked, userFirstName, userLastName, userEmail)

		firstNames := getfirstNames()

		fmt.Printf("These are first names of all the booking: %v\n ", firstNames)

		wg.Add(1)
		go sendTickets(ticketsBooked, userFirstName, userLastName, userEmail)

		if remainingTickets == 0 {
			fmt.Print("Tickets are sold. Come back next year")
		}
	} else {
		if !isValidEmail {
			fmt.Printf("Your input Email: %v is invalid\n", userEmail)
		}
		if !isValidName {
			fmt.Printf("Your input Name: %v %v is invalid\n", userFirstName, userLastName)
		}
		if !isValidTicketNumber {
			fmt.Printf("Your selected %v tickets which is more than remaining tickets: %v \n", ticketsBooked, remainingTickets)
		}

	}
	wg.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets. Remaining tickets are %v. \n", conferenceTickets, remainingTickets)
	fmt.Printf("Book your tickets to attend the %v\n", conferenceName)

}

func getfirstNames() []string {
	firstNames := []string{}
	for _, value := range bookings {
		firstNames = append(firstNames, value.firstName)
	}
	return firstNames
}

func userInput(userFirstName string, userLastName string, userEmail string, ticketsBooked uint) (string, string, string, uint) {
	fmt.Print("Please enter Your First Name: \n")
	fmt.Scan(&userFirstName)
	fmt.Print("Please enter Your Last Name: \n")
	fmt.Scan(&userLastName)
	fmt.Print("Please enter Your Email: \n")
	fmt.Scan(&userEmail)
	userFirstName = strings.Title(userFirstName)
	fmt.Printf("Welcome %v!\n", userFirstName)
	fmt.Println("Please enter number of tickets to be booked")
	fmt.Scan(&ticketsBooked)

	return userFirstName, userLastName, userEmail, ticketsBooked
}

func bookTickets(ticketsBooked uint, userFirstName string, userLastName string, userEmail string) {
	remainingTickets -= ticketsBooked

	var userData = UserData{
		firstName:     userFirstName,
		lastName:      userLastName,
		email:         userEmail,
		ticketsBooked: ticketsBooked,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of booking is %v \n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive confirmation email at %s\n", userFirstName, ticketsBooked, userEmail)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTickets(ticketsBooked uint, userFirstName string, userLastName string, userEmail string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets booked by the %v %v", ticketsBooked, userFirstName, userLastName)
	fmt.Println("#############################")
	fmt.Printf("Sending Tickets:\n %v \n to the email: %v \n", ticket, userEmail)
	fmt.Println("#############################")
	wg.Done()
}
