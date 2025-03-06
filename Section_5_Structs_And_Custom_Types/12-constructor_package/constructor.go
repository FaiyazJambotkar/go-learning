package main

func main() {
	userFirstName := GetUserData("Enter your firstname: ")
	userLastName := GetUserData("Enter your lastname: ")
	userBirthdate := GetUserData("Enter your birthdate: ")

	appUser := NewUser(userFirstName, userLastName, userBirthdate)
	appUser.OutputUserDetails()
}