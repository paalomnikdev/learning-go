package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func (u User) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
