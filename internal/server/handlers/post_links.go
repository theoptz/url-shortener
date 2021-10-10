package handlers

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/theoptz/url-shortener/internal/interfaces/iservices"
	"github.com/theoptz/url-shortener/internal/server/models"
	"github.com/theoptz/url-shortener/internal/server/restapi/operations"
)

type PostLinksHandler struct {
	links     iservices.LinksService
	validator iservices.URLValidatorService
}

func (p *PostLinksHandler) Handle(params operations.PostLinksParams) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	url := *params.Link.URL
	if !p.validator.Validate(url) {
		return operations.NewPostLinksBadRequest()
	}

	code, err := p.links.Create(ctx, url)
	if err != nil {
		return operations.NewPostLinksInternalServerError()
	}

	return operations.NewPostLinksOK().WithPayload(&models.Code{
		Code: &code,
	})
}

func NewPostLinksHandler(links iservices.LinksService, validator iservices.URLValidatorService) *PostLinksHandler {
	return &PostLinksHandler{
		links:     links,
		validator: validator,
	}
}
