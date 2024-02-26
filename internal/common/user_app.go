package common

import "errors"

const FIREBASE_APP = "firebase"
const SMS_APP = "sms"

type UserAppRequest struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
	UserUUID   string `json:"user_uuid"`
}

func (d *UserAppRequest) Validate() error {
	if d.Identifier == "" || d.UserUUID == "" || (d.Type != FIREBASE_APP && d.Type != SMS_APP) {
		return errors.New("invalid user_app")
	}
	return nil
}
