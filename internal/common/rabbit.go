package common

import "encoding/json"

// exchanges and queues
const DirectExchangeType = "direct"
const DelayedExchangeType = "x-delayed-message"

const RabbitFirstQueue = "first_queue"
const RabbitSecondQueue = "second_queue"
const ScheduledNotificationsQueue = "scheduled_notifications"

var QueueExchangeType map[string]string = map[string]string{
	RabbitFirstQueue:            DirectExchangeType,
	RabbitSecondQueue:           DelayedExchangeType,
	ScheduledNotificationsQueue: DirectExchangeType,
}

// events
type OneEvent struct {
	TicketID  int `json:"ticketId"`
	CommentID int `json:"commentId"`
}

type TwoEvent struct {
	AuthorId int `json:"authorId"`
}

type ScheduledNotification struct {
	CampaignUuid  string `json:"campaignUuid"`
	Message       string `json:"message"`
	UserUuid      string `json:"userUuid"`
	AppIdentifier string `json:"appIdentifier"`
	AppType       string `json:"appType"`
}

func NewScheduledNotification(campaignUuid, message, userUuid, appIdentifier, appType string) ([]byte, error) {
	event := &ScheduledNotification{
		CampaignUuid:  campaignUuid,
		Message:       message,
		UserUuid:      userUuid,
		AppIdentifier: appIdentifier,
		AppType:       appType,
	}

	return json.Marshal(event)
}
