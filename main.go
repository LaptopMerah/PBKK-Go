package main

import (
	"fmt"
    "strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

        isValidName := len(firstName) >= 2 && len(lastName) >= 2
        isvalidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
        isValidTickets := userTickets > 0 && userTickets <= remainingTickets

        if isValidName && isvalidEmail && isValidTickets {
            remainingTickets -= userTickets
            bookings = append(bookings, firstName+" "+lastName)
    
            fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
            fmt.Printf("We have %v tickets remaining\n", remainingTickets)
    
            firstNames := []string{}
            for _, booking := range bookings {
                var names = strings.Fields(booking)
                firstNames = append(firstNames, names[0])
            }
    
            fmt.Printf("These first names of bookings are: %v\n", firstNames)
    
            if remainingTickets == 0 {
                fmt.Println("All tickets are sold out!")
                break
            }
        } else {
            if !isValidName {
                fmt.Println("Invalid name. Name should be at least 2 characters long")
            }
            if !isvalidEmail {
                fmt.Println("Invalid email. Email should contain @ and .")
            }
            if !isValidTickets {
                fmt.Println("Invalid tickets. Please enter a valid number of tickets")
            }
        }
	}
}
