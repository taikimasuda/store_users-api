package users

import (
	"github.com/taikimasuda/store_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id 		int64 	`json:"id"`
	FirstName 	string 	`json:"first_name"`
	LastName 	string 	`json:"last_name"`
	Email 		string 	`json:"email"`
<<<<<<< HEAD
	DateCreated string 	`json:"date_created"`
	Status		string	`json:"status"`
	Password	string	`json:"-"`
=======
	DateCreated 	string 	`json:"date_created"`
>>>>>>> 1868d509564b444bd0db5fd9e7c357ec82db9298
}

type Users []User

func (user *User)Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
