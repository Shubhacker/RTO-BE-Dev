package main

import (
	"log"
	"net"

	endpoints "github.com/Shubhacker/RTO-BE-Dev/Endpoints"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Running Main")
	net, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Starting server")
	s := grpc.NewServer()

	endpoints.Endpoints(s)

	err2 := s.Serve(net)
	if err2 != nil {
		log.Println(err2.Error())
	}

}
