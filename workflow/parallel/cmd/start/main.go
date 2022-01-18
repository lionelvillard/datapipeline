package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/google/uuid"
)

func handle(e event.Event) (*event.Event, protocol.Result) {
	id := uuid.New().String()
	e.SetExtension("knativeflowid", id) // made up extension
	log.Printf("starting workflow %s\n", id)

	return &e, nil
}

func main() {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Fatalf("failed to start receiver: %s", c.StartReceiver(context.Background(), handle))
}
