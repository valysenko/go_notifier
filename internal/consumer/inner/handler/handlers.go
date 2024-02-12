package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go_notifier/internal/common"
)

type FirstQueueMessageHandler struct {
}

func (mh *FirstQueueMessageHandler) Handle(ctx context.Context, b []byte) error {
	fmt.Println("first handler")
	var event common.OneEvent
	err := json.Unmarshal(b, &event)
	if err != nil {
		return err
	}

	fmt.Println(event)
	return nil
}

type SecondQueueMessageHandler struct {
}

func (mh *SecondQueueMessageHandler) Handle(ctx context.Context, b []byte) error {
	fmt.Println("second handler")
	var event common.TwoEvent
	err := json.Unmarshal(b, &event)
	if err != nil {
		return err
	}

	fmt.Println(event)
	return nil
}
