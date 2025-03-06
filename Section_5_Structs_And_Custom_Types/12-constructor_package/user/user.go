package user

import ()
	"fmt"
	
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func NewUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func GetUserData(text string) string {
	var value string
	fmt.Print(text)
	fmt.Scanln(&value)
	return value
}

func (u *user) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}
