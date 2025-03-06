package main

import (
	"errors"
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

func newUser(firstName, lastName, birthdate string) (*user, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("values cannot be empty")
	}

	return &user {
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		creationTime: time.Now(),
	}, nil
}

func getUserData(text string) string {
	var userInputValue string
	fmt.Print(text)
	fmt.Scanln(&userInputValue)
	return userInputValue
}

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate: ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()

}

