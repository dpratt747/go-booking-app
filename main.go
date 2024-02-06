package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50

var wg = sync.WaitGroup{}

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// var bookings = []map[string]string{}
var bookings = []UserData{}
// var bookings = make([]map[string]string, 1) //alternative way to create an empty slice(list)

func main() {

	// fmt.Printf("conferenceTickets is %T, conferenceName is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)


	for {

		greetUser()

		firstName, lastName, email, userTickets := getUserInput()

		isValidEmail, isValidName, isValidUserTickers :=  helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)



		if isValidEmail && isValidName && isValidUserTickers {

			wg.Add(1)
			bookTickets(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			printFirstNames()

			if remainingTickets == 0 {
				fmt.Println("Conference tickets are all sold out")
				break
			}

		} else {
			
			var messages []string

			if (!isValidEmail) {
				messages = append(messages, "Your email address does not contain an @ symbol")
			} 
			
			if (!isValidName) {
				messages = append(messages, "Your name is incorrect both the first name and the last name need to be greater than two characters")
			}
			
			if (!isValidUserTickers) {
				messages = append(messages, "You requested an invalid number of tickets; the amount requested must be greater than 0 and less than or equal to the remaining number of tickets")
			}

			for _, message := range messages {
				fmt.Println(message)
			}

		}


		// city := "London"

		// switch strings.ToLower(city) {
		// 	case "london":
		// 		fmt.Println("London has been matched")
		// 	default:
		// 		fmt.Println("City provided does not match expected")
		// }

	}

	wg.Wait()
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// userData := make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	userData := UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	// fmt.Printf("The whole bookings slice: %v\n", bookings)
	// fmt.Printf("The first value in the bookings slice: %v\n", bookings[0])
	// fmt.Printf("The bookings slice is %v long.\n", len(bookings))
	// fmt.Printf("The type of the bookings slice: %T\n", bookings)

	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames(){
	firstNames := []string {}

	for _, userDataStruct := range bookings {
		firstNames = append(firstNames, userDataStruct.firstName)
	}

	fmt.Printf("These are the first names %v\n", firstNames)
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
	
	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string){
	time.Sleep(50 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)

	fmt.Println("------------------------------------------------")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, emailAddress)
	fmt.Println("------------------------------------------------")
	wg.Done()
}