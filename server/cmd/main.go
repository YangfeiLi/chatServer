package main

import (
	"log"

	"google.golang.org/grpc"

	"../internal"
	"../pkg/api"
)

func main() {
	log.Println("Launching message server...")

	portbinder := &internal.PortBinder{
		Port: 5558,
	}

	server := grpc.NewServer()
	go api.RegisterMessagingServer(server, &internal.MessagerServer{})

	portbinder.Bind(server)
}
