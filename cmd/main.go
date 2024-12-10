package main

import (
	"booking-app/pkg/common"
	"fmt"
	"sync"
	"time"
)

var (
	//conference data
	conferenceName string = "Go Converence"
	conferenceRemainingTickets uint = common.ConferenceTickets
	
	//user booking data
	bookings = make([]UserData, 0)
)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	//for {
		
		userfirstName,userlastName,userEmail,userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := common.ValidateUserInput(userfirstName, userlastName, userEmail, userTickets, conferenceRemainingTickets)

		fmt.Println("Processing your ticket...")	
		if isValidName && isValidEmail && isValidTickets {
			
			bookTicket(userfirstName, userlastName, userEmail, userTickets)

			wg.Add(1)
			go sendTicket(userfirstName, userlastName, userEmail, userTickets)
			
			firstNames := getFirstNames()
			fmt.Printf("Current booking attendees(first name) are: %s\n", firstNames)


			if conferenceRemainingTickets == 0 {
				fmt.Println("Our Conference is sold out!. Come again next time!")
				//break
			}
		} else {
			fmt.Printf("Your data input is invalid, please try again \n")
			if !isValidName {
				fmt.Println("Invalid name, name must be at least 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Invalid email, email must contain @")
			}
			if !isValidTickets {
				fmt.Println("Invalid ticket, ticket must be greater than 0 and less than or equal to remaining tickets")
			}
		}
		wg.Wait()
	//}
}

func greetUser () {
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("we currently have a total of %d but only %d tickets left!\n", common.ConferenceTickets, conferenceRemainingTickets)
	fmt.Println("Get your attending tickets here!")
	
}

func getFirstNames () []string {
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	} 
	return firstNames
			
}



func getUserInput() (string, string, string, uint) {
	//user data
	var (
	userfirstName string
	userlastName string
	userEmail string
	userTickets uint
	)
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&userfirstName)
	
	fmt.Println("Enter your lastname: ")
	fmt.Scan(&userlastName)
	
	fmt.Println("Enter your email: ")
	fmt.Scan(&userEmail)
	
	fmt.Println("Enter number of ticket: ")
	fmt.Scan(&userTickets)

	return userfirstName, userlastName, userEmail, userTickets
}
func bookTicket (userfirstName string, userlastName string, userEmail string, userTickets uint) {

	conferenceRemainingTickets -= userTickets

	var userData = UserData{
		firstName: userfirstName,
		lastName: userlastName,
		email: userEmail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings : %v\n", bookings)
	
	fmt.Printf("Thank you %s %s for booking %d tickets to %s. Your ticket has been sent to %s\n", userfirstName,userlastName, userTickets, conferenceName, userEmail)
	fmt.Printf("We now have %d tickets left\n", conferenceRemainingTickets)
}

func sendTicket( firstName string, lastName string, email string, userTickets uint){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%d number of tickets for %s %s", userTickets, firstName, lastName)
	fmt.Println("------------------------------------------------")
	fmt.Printf("Sending ticket :\n %v\n to email address %s\n", ticket, email)
	fmt.Println("------------------------------------------------")
	wg.Done()
}