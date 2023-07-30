package repository

import (
	"context"

	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
)

type APIPorts interface {
	GetCreate(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error)
}
