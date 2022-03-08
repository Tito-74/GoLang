package main

import (
	"fmt"
	"time"
	"sync"
	
)

const conferenceTicket = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct{
	firstName string
	lastName string
	email string
	userTickets uint

}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// var bookings []string


	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidemail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)

	if isValidName && isValidemail && isValidTicketNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)

		go sendingTicket(userTickets, firstName, lastName, email)

		fmt.Printf("All bookings are : %v \n", bookings)

		firstNames := getFirstName()
		fmt.Printf("All bookings are : %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out . Come back next time \n")
			// break
		}

	} else {
		// fmt.Printf("You input is incorrect \n")
		if !isValidName {
			fmt.Printf("Please enter the correct Name \n")
		}
		if !isValidemail {
			fmt.Printf("Please enter the correct Email \n")
		}

		if !isValidTicketNumber {
			fmt.Printf("Please enter the correct Ticket Number \n")
		}

	}
	wg.Wait()



}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n ", conferenceName)
	fmt.Printf("We have total number of Tickets %v and the no of tickets remaining are %v \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you need:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	
	// Creating a map
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,

	}


	bookings = append(bookings, userData)

	fmt.Println("The first value: \n", bookings[0])

	fmt.Printf("%v ,thank you for  booking %v tickets With us you will recieve a confirmation email on %v \n", firstName, userTickets, email)

	fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)
}


func sendingTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n",userTickets,firstName, lastName )
	fmt.Printf("#######################################\n")
	fmt.Printf("sending %v \nto email address %v \n ",ticket, email)
	fmt.Printf("#######################################\n")
	wg.Done()

}
