package main

import (
	"context"
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"
)

func handle(e event.Event) (*event.Event, protocol.Result) {
	id, err := e.Context.GetExtension("knativeflowid")
	if err != nil {
		return nil, protocol.NewReceipt(true, "missing knativeflowid attribute")
	}

	log.Printf("short delay for %s\n", id)
	time.Sleep(5 * time.Second)

	return &e, protocol.ResultACK
}

func main() {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Fatalf("failed to start receiver: %s", c.StartReceiver(context.Background(), handle))
}
