package controller

import (
	"context"

	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
)

type EndpointsInterface interface {
	Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error)
}
