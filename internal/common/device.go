package common

import "errors"

type DeviceRequest struct {
	Token    string `json:"token"`
	UserUUID string `json:"user_uuid"`
}

func (d *DeviceRequest) Validate() error {
	if d.Token == "" || d.UserUUID == "" {
		return errors.New("invalid device")
	}
	return nil
}
