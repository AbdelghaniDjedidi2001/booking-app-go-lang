package main

import "fmt"

func main() {
	// var name = "Go Confrerence"
	const confrerenceName = "Go Confrerence"
	const confrerenceTickets = 50
	remainingTickets := 50 //same as  var remainingTickets = 50 and work only in var not const 

	fmt.Printf("confrerenceName is %T, confrerenceTickets is %T, remainingTickets is %T\n",confrerenceName,confrerenceTickets,remainingTickets)

	// fmt.Println("Welcome to our ", name, "booking application")
	// fmt.Println("we have total of", confrerenceTickets, "tickets and ", remainingTickets, "are still available.")

	fmt.Printf("Welcome to our %s booking application\n",confrerenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n",confrerenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var userName error
	var firstName string
	var lastName string
	var email string
	var userTickets int



	// fmt.Println(&userName)
	// fmt.Println(&remainingTickets)
	fmt.Print("Enter your firstName : ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your lastName : ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email : ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets : ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %s %s for booking %v tickets, you will receive a confermation email  at %s\n",firstName,lastName,userTickets,email)
	
}
