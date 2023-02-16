package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ZhijiunY/booking-app/helper"
)

const conferenceTickets int = 50

var conferenceName string = "Go conference"
var remainingTickets uint = 50

// create a empty of list of maps
// 0 is initial size or 1 can be initial size too
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!", conferenceName)
	// fmt.Printf("conferencesTickets is %T, conferenceName is %T, remainingTickets is %T\n", conferenceTickets, conferenceName, remainingTickets)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("You can get your tickets herer to attend.")

}

// func printFirstName(something []input value)[]out value{}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	// map[key]value
	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// // userData["key"] = value
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// // base 10, is decimal numbers
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings[0] = firstName + "" + lastName
	// firstName+" "+lastName change to userData
	bookings = append(bookings, userData)
	fmt.Printf("List of Data is: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("================================================================")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("================================================================")
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	// validate
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		// call function printFirstName
		firstNames := getFirstNames()
		fmt.Printf("The first name of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out. Come back next year.")
			//break
		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
		fmt.Println("Your input data is invalid. Please try again.")
	}
	wg.Wait()

}
