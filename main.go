package main

import "fmt"

func main() {

	conferenceName := "go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("welcome to  %v booking application\n ", conferenceName)
	fmt.Printf("we have total of %v tickets remainings = %v\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")

}
