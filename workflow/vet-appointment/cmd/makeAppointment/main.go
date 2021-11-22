package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"
)

func handle(e event.Event) (*event.Event, protocol.Result) {
	id, err := e.Context.GetExtension("knativeflowid")
	if err != nil {
		return nil, protocol.NewReceipt(true, "missing knativeflowid attribute")
	}

	log.Printf("making appointment for %s\n", id)

	er := event.New()
	er.SetID(e.ID())
	er.SetSource("VetServiceSource")
	er.SetType("events.vet.appointments")
	er.SetExtension("knativeflowid", id)
	err = er.SetData("application/json", `{"appointmentInfo": {"time": "tuesday"}}`)
	if err != nil {
		return nil, protocol.NewReceipt(true, "internal error: %v", err)
	}

	return &er, protocol.ResultACK
}

func main() {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Fatalf("failed to start receiver: %s", c.StartReceiver(context.Background(), handle))
}
