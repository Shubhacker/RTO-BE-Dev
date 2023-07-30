package repository

import (
	"context"
	"log"

	"github.com/Shubhacker/RTO-BE-Dev/controller"
	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
)

type Adapters struct {
	cr controller.EndpointsInterface
}

func NewAdapter(apiport controller.EndpointsInterface) *Adapters {
	return &Adapters{
		cr: apiport,
	}
}

func (api Adapters) GetCreate(ctx context.Context, i *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error) {
	var result *invoicer.CreateResponseMessage
	log.Println("inside Adapter package")
	api.cr.Create(ctx, i)
	return result, nil
}
