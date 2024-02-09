package common

import "errors"

type CampaignRequest struct {
	Name       string   `json:"name"`
	Message    string   `json:"message"`
	Time       string   `json:"time"`
	DaysOfWeek []string `json:"days_of_week"`
}

func (c *CampaignRequest) Validate() error {
	if c.Name == "" || c.Message == "" || c.Time == "" || len(c.DaysOfWeek) == 0 {
		return errors.New("invalid campaign")
	}
	return nil
}
