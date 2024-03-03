package notifier

import "context"

const FIREBASE_NOTIFIER_TYPE = "firebase"

type Firebase interface {
	SendNotification(ctx context.Context, toDeviceToken, message string) error
}

type FirebaseNotifier struct {
	firebase Firebase
}

func NewFirebaseNotifier(f Firebase) *FirebaseNotifier {
	return &FirebaseNotifier{
		firebase: f,
	}
}

func (n *FirebaseNotifier) Supports(appType string) bool {
	return appType == FIREBASE_NOTIFIER_TYPE
}

func (n *FirebaseNotifier) Notify(ctx context.Context, appIdentifier, message string) error {
	return n.firebase.SendNotification(ctx, appIdentifier, message)
}
