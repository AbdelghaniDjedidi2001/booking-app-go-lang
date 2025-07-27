package main

import (
	"fmt"
	"strings"
)

func main() {
	// var name = "Go Confrerence"
	const confrerenceName = "Go Confrerence"
	const confrerenceTickets = 50
	// remainingTickets := 50 //same as  var remainingTickets = 50 and work only in var not const
	var remainingTickets uint = 50

	fmt.Printf("confrerenceName is %T, confrerenceTickets is %T, remainingTickets is %T\n", confrerenceName, confrerenceTickets, remainingTickets)

	// fmt.Println("Welcome to our ", name, "booking application")
	// fmt.Println("we have total of", confrerenceTickets, "tickets and ", remainingTickets, "are still available.")

	fmt.Printf("Welcome to our %s booking application\n", confrerenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", confrerenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings [50]string    array
	var bookings []string // slice
	// var index uint = 0

	// var userName error
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// for index < 50 && remainingTickets != 0 {
	for remainingTickets != 0 {
		// fmt.Println(&userName)
		// fmt.Println(&remainingTickets)
		fmt.Printf("Remaining Tickets : %v\n", remainingTickets)

		var isFirstNameValid = false
		for !isFirstNameValid {
			fmt.Print("Enter your firstName : ")
			fmt.Scan(&firstName)
			isFirstNameValid = len(firstName) > 2
		}

		var isLastNameValid = false
		for !isLastNameValid {
			fmt.Print("Enter your lastName : ")
			fmt.Scan(&lastName)
			isLastNameValid = len(lastName) > 2
		}

		var isEmailValid = false
		for !isEmailValid {
			fmt.Print("Enter your email : ")
			fmt.Scan(&email)
			isEmailValid = strings.Contains(email, "@")
		}

		fmt.Print("Enter number of tickets : ")
		fmt.Scan(&userTickets)

		for userTickets > remainingTickets || userTickets == 0 {
			fmt.Printf("Enter number of tickets less than %v: ", remainingTickets)
			fmt.Scan(&userTickets)
		}

		remainingTickets -= userTickets
		// var userData = "full Name %s %s,Number of tickets %v,Email %s" ,firstName,lastName,userTickets,email
		userData := fmt.Sprintf("full Name %s %s, Number of tickets %v, Email %s", firstName, lastName, userTickets, email)
		// bookings[index] = userData   array
		bookings = append(bookings, userData) //slice

		fmt.Printf("Thank you %s %s for booking %v tickets, you will receive a confermation email  at %s\n", firstName, lastName, userTickets, email)

		// index++
	}

	fmt.Printf("all booking : %v\n", bookings)
}
