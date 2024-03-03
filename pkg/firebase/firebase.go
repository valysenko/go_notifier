package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

// https://firebase.google.com/docs/admin/setup/#set-up-project-and-service-account
// https://firebase.google.com/docs/cloud-messaging/send-message
// https://blog.canopas.com/golang-send-firebase-push-notifications-to-apps-45bc83d7cabf
type Firebase struct {
	AuthConfig []byte
	logger     log.FieldLogger
}

func NewFirebase(config []byte, logger log.FieldLogger) *Firebase {
	return &Firebase{
		AuthConfig: config,
		logger:     logger,
	}
}

func (f *Firebase) SendNotification(ctx context.Context, toDeviceToken, message string) error {
	opts := []option.ClientOption{option.WithCredentialsJSON(f.AuthConfig)}
	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		return err
	}

	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		return err
	}

	response, err := fcmClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: "App notification",
			Body:  message,
		},
		Token: toDeviceToken,
	})
	if err != nil {
		return err
	}

	f.logger.Info("sent message to firebase. response: " + response)

	return nil
}

// many
// response, err := fcmClient.SendEach(context.Background(), []*messaging.Message{&messaging.Message{
// 	Notification: &messaging.Notification{
// 		Title: "Congratulations01!!",
// 		Body:  "You have just implement push notification",
// 	},
// 	Token: token,
// }, &messaging.Message{
// 	Notification: &messaging.Notification{
// 		Title: "Congratulations02!!",
// 		Body:  "You have just implement push notification",
// 	},
// 	Token: token,
// }})
