package service

import (
	"context"
	"log"

	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
)

type Adapters struct {
}

func NewAdapter() *Adapters {
	return &Adapters{}
}

func (A Adapters) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error) {
	log.Println("Inside service")
	var result *invoicer.CreateResponseMessage
	return result, nil
}
