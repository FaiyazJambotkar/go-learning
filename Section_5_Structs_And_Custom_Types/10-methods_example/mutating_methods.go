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

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserDetails() {
	u.lastName = ""
	u.birthdate = ""
}

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate: ")

	appUser := user {
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		creationTime: time.Now(),
	}

	appUser.outputUserDetails()
	appUser.clearUserDetails()
	appUser.outputUserDetails()
}

func getUserData(text string) string {
	var userInputValue string
	fmt.Print(text)
	fmt.Scan(&userInputValue)
	return userInputValue
}

