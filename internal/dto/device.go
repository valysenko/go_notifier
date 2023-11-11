package dto

import "errors"

type Device struct {
	Token    string `json:"token"`
	UserUUID string `json:"user_uuid"`
}

func (d *Device) Validate() error {
	if d.Token == "" || d.UserUUID == "" {
		return errors.New("invalid device")
	}
	return nil
}
