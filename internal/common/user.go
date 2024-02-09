package common

import "errors"

type UserRequest struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Timezone string `json:"timezone"`
}

func (u *UserRequest) Validate() error {
	if u.UUID == "" || u.Email == "" || u.Timezone == "" {
		return errors.New("invalid user")
	}

	return nil
}
