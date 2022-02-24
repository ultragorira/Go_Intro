package main

import "fmt"

func main() {
	confName := "Go Conference"
	const confTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our system.\nReserve your ticket for the %v\n", confName)
	fmt.Println("There is a total of", confTickets, "and left are", remainingTickets)
	var firstName string
	var lastName string
	var email string
	var numTickets uint

	fmt.Println("Give your name:")
	fmt.Scan(&firstName)

	fmt.Println("Give your surname:")
	fmt.Scan(&lastName)

	fmt.Println("Give your email address:")
	fmt.Scan(&email)

	fmt.Println("How many tickets:")
	fmt.Scan(&numTickets)

	remainingTickets = remainingTickets - numTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. Email sent to %v", firstName, lastName, numTickets, email)
	fmt.Printf("Tickets lefts are %v", remainingTickets)
}
