package main

import "fmt"

func main() {
	var conferenceNmae="GO Conference"
	const conferenceTickets=50
	var remainingTickets=50

	fmt.Printf("Conference Name type is '%T' conference ticket type is '%T' \n", conferenceNmae, conferenceTickets)

	fmt.Println("Hello Welcome to", conferenceNmae, "!")
	fmt.Println("We have overall", conferenceTickets, "and", remainingTickets,  "number of Tickets are available for now")
	fmt.Println("Get your tickets here to attend")
    
	var firstName string
	var LastName string
	var email string
	var age int
	var userTickets int
	var isTrue bool

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&LastName)
	fmt.Println("Please enter your age:")
	fmt.Scan(&age)
	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Please enter your how number of tickets you want to book:")
	fmt.Scan(&userTickets)
	fmt.Println("Please confirm your order: True/False")
	fmt.Scan(&isTrue)

	if isTrue {
		fmt.Printf("Confirmation:\n First Name: %v \n Last Name: %v \n Age: %v \n Email: %v \n Number of Tickets: %v \n", firstName, LastName, age, email, userTickets)
		remainingTickets=remainingTickets-userTickets
		fmt.Printf("Remaining ticktes after booking: %v", remainingTickets)
	} else {
		fmt.Print("You didn't confirm! Please fill out again")
	}

	



	

}