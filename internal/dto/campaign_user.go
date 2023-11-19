package dto

import "errors"

type CampaignUser struct {
	CampaignUUID string `json:"campaign_uuid"`
	UserUUID     string `json:"user_uuid"`
}

func (d *CampaignUser) Validate() error {
	if d.CampaignUUID == "" || d.UserUUID == "" {
		return errors.New("invalid campaign_user")
	}
	return nil
}
