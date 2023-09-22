package main

import (
	"log"
	"net"

	DB "github.com/Shubhacker/RTO-BE-Dev/Adapters/Database"
	"github.com/Shubhacker/RTO-BE-Dev/Adapters/model"
	endpoints "github.com/Shubhacker/RTO-BE-Dev/Endpoints"
	"google.golang.org/grpc"
)

var DBConn *model.Conn

func main() {
	log.Println("Running Main")
	net, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Starting server")
	s := grpc.NewServer()
	DBConn = DB.DBInit()
	endpoints.Endpoints(s)

	err2 := s.Serve(net)
	if err2 != nil {
		log.Println(err2.Error())
	}

}
