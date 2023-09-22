package endpoints

import (
	"context"
	"log"
	"time"

	"github.com/Shubhacker/RTO-BE-Dev/controller/service"
	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s MyInvoicerServer) ReportSocialComments(ctx context.Context, i *invoicer.ReportSocialCommentsRequest) (*invoicer.ReportSocialCommentsResponse, error) {
	// var response invoicer.ReactSocialResponce
	service := service.NewAdapter()
	response, err := service.ReportSocialCommentsService(ctx, i)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s MyInvoicerServer) ReactOnSocialReport(ctx context.Context, i *invoicer.ReactSocialRequest) (*invoicer.ReactSocialResponce, error) {
	log.Println("Inside React Social Post API")
	now := time.Now()
	service := service.NewAdapter()
	response, err := service.ReactSocialPost(ctx, i)
	if err != nil {
		return response, err
	}
	log.Println(time.Since(now), "<-> Time took for API")
	return response, nil
}

func (s MyInvoicerServer) Create(ctx context.Context, i *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error) {
	log.Println("Inside Endpoint")
	service := service.NewAdapter()
	service.Create(ctx, i)
	return &invoicer.CreateResponseMessage{
		Pdf: []byte("Test"),
	}, nil
}

func (s MyInvoicerServer) CreateReport(ctx context.Context, i *invoicer.CreateReportRequest) (*invoicer.CreateReportResponse, error) {
	// var response invoicer.CreateReportResponse
	log.Println("Inside Endpoint")
	now := time.Now()
	service := service.NewAdapter()
	response, err := service.CreateReport(ctx, i)
	if err != nil {
		return nil, err
	}
	log.Println(time.Since(now), "<-> Time took for API")
	return response, nil
}

func (s MyInvoicerServer) Social(ctx context.Context, i *invoicer.SocialRequest) (*invoicer.SocialData, error) {
	// var response invoicer.CreateReportResponse
	now := time.Now()
	log.Println("Inside Endpoint")
	service := service.NewAdapter()
	response, err := service.Social(ctx, i)
	if err != nil {
		return nil, err
	}
	log.Println(time.Since(now), "<-> Time took for API")
	return response, nil
}

func Endpoints(s grpc.ServiceRegistrar) {
	service := &MyInvoicerServer{}

	invoicer.RegisterInvoicerServer(s, service)
}
