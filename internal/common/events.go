package common

import "encoding/json"

const RabbitFirstQueue = "first_queue"
const RabbitSecondQueue = "second_queue"
const ScheduledNotificationsQueue = "scheduled_notifications"

type OneEvent struct {
	TicketID  int `json:"ticketId"`
	CommentID int `json:"commentId"`
}

type TwoEvent struct {
	AuthorId int `json:"authorId"`
}

type ScheduledNotification struct {
	CampaignUuid string `json:"campaignUuid"`
	Message      string `json:"message"`
	UserUuid     string `json:"userUuid"`
	DeviceToken  string `json:"deviceToken"`
}

func NewScheduledNotification(campaignUuid, message, userUuid, deviceToken string) ([]byte, error) {
	event := &ScheduledNotification{
		CampaignUuid: campaignUuid,
		Message:      message,
		UserUuid:     userUuid,
		DeviceToken:  deviceToken,
	}

	return json.Marshal(event)
}
