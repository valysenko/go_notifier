package common

import "errors"

type CampaignUserRequest struct {
	CampaignUUID string `json:"campaign_uuid"`
	UserUUID     string `json:"user_uuid"`
}

func (d *CampaignUserRequest) Validate() error {
	if d.CampaignUUID == "" || d.UserUUID == "" {
		return errors.New("invalid campaign_user")
	}
	return nil
}
