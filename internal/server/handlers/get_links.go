package handlers

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/theoptz/url-shortener/internal/interfaces/iservices"
	"github.com/theoptz/url-shortener/internal/server/models"
	"github.com/theoptz/url-shortener/internal/server/restapi/operations"
)

type GetLinksHandler struct {
	links iservices.LinksService
}

func (g *GetLinksHandler) Handle(params operations.GetLinksCodeParams) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	link, err := g.links.GetByCode(ctx, params.Code)
	if err != nil {
		return operations.NewGetLinksCodeInternalServerError()
	} else if link == "" {
		return operations.NewGetLinksCodeNotFound()
	}

	return operations.NewGetLinksCodeOK().WithPayload(&models.Link{URL: &link})
}

// NewGetLinksHandler returns *GetLinksHandler pointer
func NewGetLinksHandler(links iservices.LinksService) *GetLinksHandler {
	return &GetLinksHandler{
		links: links,
	}
}
