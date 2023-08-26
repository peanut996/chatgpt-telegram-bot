package model

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	UserName string `json:"username,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`
}

func (u *User) String() string {
	if u.UserName != "" {
		return fmt.Sprintf("@%s", u.UserName)
	}
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func NewUser(username, firstName, lastName string) *User {
	return &User{
		UserName:  username,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func From(u *tgbotapi.User) *User {
	return &User{
		UserName:  u.UserName,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
