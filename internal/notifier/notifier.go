package notifier

type Notifier interface {
	Notify(identifier, message string) error
}

type FirebaseNotifier struct {
}

type TwillioNotifier struct {
}
