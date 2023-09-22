package controller

import (
	"context"

	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
)

type EndpointsInterface interface {
	Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error)
	CreateReport(ctx context.Context, i *invoicer.CreateReportRequest) (*invoicer.CreateReportResponse, error)
	Social(ctx context.Context, i *invoicer.SocialRequest) (*invoicer.SocialData, error)
	ReportSocialComments(ctx context.Context, i *invoicer.ReportSocialCommentsRequest) (*invoicer.ReactSocialResponce, error)
}
