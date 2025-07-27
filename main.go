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
	var userName string
	var userTickets int

	// userName = "Ghonou"
	userTickets = 2

	// fmt.Println(&userName)
	// fmt.Println(&remainingTickets)

	fmt.Scan(&userName)

	fmt.Printf("User %s booked %v tickets\n",userName,userTickets)
	
}
