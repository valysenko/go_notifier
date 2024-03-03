package handler

import (
	"context"
	"errors"
	"fmt"
	"go_notifier/internal/app/consumer/inner/handler/mocks"
	"go_notifier/internal/common"
	notifierMocks "go_notifier/internal/notifier/mocks"
	"go_notifier/pkg/transport/rabbitmq"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	log.SetLevel(log.FatalLevel)
	appId := "appId"
	appType := common.FIREBASE_APP
	msg := "msg"
	e, _ := common.NewScheduledNotification("cUuid", msg, "uUuid", appId, appType)
	nProvider := mocks.NewNotifierProvider(t)
	handler := NewScheduledNotificationHandler(log.New(), nProvider)
	notifier := notifierMocks.NewNotifier(t)
	ctx := context.Background()

	t.Run("scheduled notification handler success", func(t *testing.T) {
		notifier.On("Notify", ctx, appId, msg).Once().Return(nil)
		nProvider.On("Provide", common.FIREBASE_APP).Once().Return(notifier)
		err := handler.Handle(ctx, e)
		assert.Nil(t, err)
	})

	t.Run("scheduled notification handler error", func(t *testing.T) {
		notifierErr := errors.New("something went wrong")
		expectedErr := rabbitmq.NewRetriableError(notifierErr, "error while processing message")
		notifier.On("Notify", ctx, appId, msg).Once().Return(notifierErr)
		nProvider.On("Provide", common.FIREBASE_APP).Once().Return(notifier)
		err := handler.Handle(ctx, e)
		assert.Equal(t, expectedErr, err)
		fmt.Println(err)
	})
}
