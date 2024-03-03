package notifier

import "context"

const SMS_NOTIFIER_TYPE = "sms"

type Twilio interface {
	SendNotification(ctx context.Context, toNumber, message string) error
}

type TwilioNotifier struct {
	twilio Twilio
}

func NewTwilioNotifier(t Twilio) *TwilioNotifier {
	return &TwilioNotifier{
		twilio: t,
	}
}

func (n *TwilioNotifier) Supports(appType string) bool {
	return appType == SMS_NOTIFIER_TYPE
}

func (n *TwilioNotifier) Notify(ctx context.Context, appIdentifier, message string) error {
	return n.twilio.SendNotification(ctx, appIdentifier, message)
}
