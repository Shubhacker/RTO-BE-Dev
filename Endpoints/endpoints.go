package endpoints

import (
	"context"
	"log"

	"github.com/Shubhacker/RTO-BE-Dev/controller/service"
	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s MyInvoicerServer) Create(ctx context.Context, i *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error) {
	log.Println("Inside Endpoint")
	service := service.NewAdapter()
	service.Create(ctx, i)
	return &invoicer.CreateResponseMessage{
		Pdf: []byte("Test"),
	}, nil
}

func Endpoints(s grpc.ServiceRegistrar) {
	service := &MyInvoicerServer{}

	invoicer.RegisterInvoicerServer(s, service)
}
