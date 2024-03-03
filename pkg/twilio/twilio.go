package twilio

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	twilioLib "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// https://www.twilio.com/en-us/blog/send-sms-30-seconds-golang
type Twilio struct {
	accountSid    string
	authToken     string
	accountNumber string
	client        *twilioLib.RestClient
	logger        log.FieldLogger
}

func NewTwilio(accountSid, authToken, accountNumber string, logger log.FieldLogger) *Twilio {
	client := twilioLib.NewRestClientWithParams(twilioLib.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &Twilio{
		accountSid:    accountSid,
		authToken:     authToken,
		accountNumber: accountNumber,
		client:        client,
		logger:        logger,
	}
}

func (t *Twilio) SendNotification(ctx context.Context, toNumber, message string) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(t.accountNumber)
	params.SetBody(message)

	resp, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	response, err := json.Marshal(*resp)
	if err != nil {
		return err
	}

	t.logger.Info("sent message to twillio. response: " + string(response))

	return nil
}
