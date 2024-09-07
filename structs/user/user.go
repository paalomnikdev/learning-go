package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin {
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
