package main

import "fmt"

func main() {
	// var name = "Go Confrerence"
	const name = "Go Confrerence"
	const confrerenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to our ", name, "booking application")
	fmt.Println("we have total of", confrerenceTickets, "tickets and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}
