package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)



const conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {


	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, userTickets, remainingTickets, email, conferenceName)

			wg.Add(1)
			go sendTicket(userTickets,firstName,lastName,email)

		

			
			if remainingTickets == 0 {
				// end the program
				fmt.Print("\nSorry, our conference is booked up! Come back next year")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("Please enter valid emailId")
			}
			if !isValidName {
				fmt.Println("Please enter valid name")
			}
			if !isValidTicketNumber {
				fmt.Println("Please enter valid Ticket count")
			}

		}
	}
	wg.Wait()
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("There are %v in total as of now, out of which, %v are available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend!")
}

// func getFirstNames() []string {
// 	firstNames := []string{}
// 	for _, booking := range bookings {		
// 		firstNames = append(firstNames, booking.firstName)
// 	}
// 	return firstNames
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("\nEnter first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("\nEnter last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("\nEnter email: ")
	fmt.Scan(&email)
	fmt.Printf("\nNo. of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
func bookTicket(firstName string, lastName string, userTickets uint, remainingTickets uint, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	// create a struct
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		noOfTickets: userTickets,
	}
	
	

	bookings = append(bookings , userData)
	fmt.Print("List of bookings:\n" , bookings)

	fmt.Printf("%v user booked %v tickets\n", firstName, userTickets)
	fmt.Printf("Confirmation will be sent at %v\n", email)
	fmt.Printf("As of now, %v tickets are remaining for %v ", remainingTickets, conferenceName)
}
func sendTicket(userTickets uint,firstName string,lastName string,emailId string){
	time.Sleep(10 * time.Second)
	var msg = fmt.Sprintf("%d tickets for %v %v",userTickets,firstName,lastName)
	fmt.Printf("******")
	fmt.Printf("Sending ticket: %v to %v", msg, emailId)
	fmt.Printf("******")
	wg.Done()
}