package main

import "fmt"

func main() {
	var conferenceName = "go conference"
	const conferenceTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking server")
	fmt.Println("You can get your tickets herer to attend")

	fmt.Println(conferenceName)
	fmt.Println(conferenceTickets)

}
