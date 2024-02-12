package common

const RabbitFirstQueue = "first_queue"
const RabbitSecondQueue = "second_queue"

type OneEvent struct {
	TicketID  int `json:"ticketId"`
	CommentID int `json:"commentId"`
}

type TwoEvent struct {
	AuthorId int `json:"authorId"`
}
