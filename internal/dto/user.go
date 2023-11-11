package dto

import "errors"

type User struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Timezone string `json:"timezone"`
}

func (u *User) Validate() error {
	if u.UUID == "" || u.Email == "" || u.Timezone == "" {
		return errors.New("invalid user")
	}

	return nil
}
