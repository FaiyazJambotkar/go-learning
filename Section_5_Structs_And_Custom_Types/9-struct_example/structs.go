package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	creationTime time.Time
}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthdate := getUserData("Enter your birthdate: ")

	// var appUser user 
	appUser := user {
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		creationTime: time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(inputText string) string {
	fmt.Print(inputText)
	var userData string
	fmt.Scan(&userData)
	return userData
}
