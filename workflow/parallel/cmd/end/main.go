package main

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/cloudevents/sdk-go/v2/types"
)

func handle(e event.Event) (*event.Event, protocol.Result) {
	id, err := types.ToString(e.Context.GetExtensions()["knativeflowid"])

	if err != nil {
		log.Println("missing or invalid knativeflowid attribute")
	} else {
		log.Printf("workflow %s terminated\n", id)
	}

	fmt.Println(e)

	return nil, protocol.ResultACK
}

func main() {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Fatalf("failed to start receiver: %s", c.StartReceiver(context.Background(), handle))
}
