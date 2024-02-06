package main

import (
	"context"
	"fmt"
	"go_notifier/pkg/rabbitmq"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	amqpConnection := rabbitmq.NewConnection()
	defer amqpConnection.Close()

	consumers := rabbitmq.NewConsumersMap(amqpConnection)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	grp, ctx := errgroup.WithContext(ctx)
	for _, consumer := range consumers {
		consumer := consumer
		grp.Go(func() error {
			return consumer.Listen(ctx)
		})
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Println("SIGTERM was send, close consumer")
		consumers.Close()
		amqpConnection.Close()
		cancel()
	}()

	if err := grp.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println("All consumers have finished processing.")
	}
}
