package notifier

import "context"

type NotifierProvider struct {
	notifiers []Notifier
}

func NewNotifierProvider(notifiers ...Notifier) *NotifierProvider {
	return &NotifierProvider{
		notifiers: notifiers,
	}
}

func (np *NotifierProvider) Provide(appType string) Notifier {
	for _, notifier := range np.notifiers {
		if notifier.Supports(appType) {
			return notifier
		}
	}

	return nil
}

// go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Notifier --case snake
type Notifier interface {
	Supports(appType string) bool
	Notify(ctx context.Context, appIdentifier, message string) error
}
