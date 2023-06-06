package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return
	}
	// defer nc.Close()

	echoHandler := func(req micro.Request) {
		req.Respond(req.Data())
	}

	config := micro.Config{
		Name:        "EchoService",
		Version:     "1.0.0",
		Description: "Send back what you receive",
		// DoneHandler can be set to customize behavior on stopping a service.
		DoneHandler: func(srv micro.Service) {
			info := srv.Info()
			fmt.Printf("stopped service %q with ID %q\n", info.Name, info.ID)
		},

		// ErrorHandler can be used to customize behavior on service execution error.
		ErrorHandler: func(srv micro.Service, err *micro.NATSError) {
			info := srv.Info()
			fmt.Printf("Service %q returned an error on subject %q: %s", info.Name, err.Subject, err.Description)
		},

		// optional base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "echo",
			Handler: micro.HandlerFunc(echoHandler),
		},
	}

	micro.AddService(nc, config)

	runtime.Goexit()
}
