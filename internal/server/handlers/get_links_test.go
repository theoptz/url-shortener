package handlers

import (
	"context"
	"errors"
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/theoptz/url-shortener/internal/interfaces/iservices"
	"github.com/theoptz/url-shortener/internal/server/models"
	"github.com/theoptz/url-shortener/internal/server/restapi/operations"
)

func TestNewGetLinksHandler(t *testing.T) {
	links := &iservices.MockLinksService{}

	h := NewGetLinksHandler(links)

	assert.Equal(t, links, h.links)
}

func TestGetLinksHandler_Handle(t *testing.T) {
	exampleLink := "https://example.com/"

	testCases := []struct {
		params operations.GetLinksCodeParams

		returnedLink string
		returnedErr  error

		res middleware.Responder
	}{
		{
			params: operations.GetLinksCodeParams{
				Code: "unknown code",
			},

			returnedErr: errors.New("internal server error"),

			res: operations.NewGetLinksCodeInternalServerError(),
		},
		{
			params: operations.GetLinksCodeParams{
				Code: "not found code",
			},

			returnedLink: "",

			res: operations.NewGetLinksCodeNotFound(),
		},
		{
			params: operations.GetLinksCodeParams{
				Code: "example code",
			},

			returnedLink: exampleLink,

			res: operations.NewGetLinksCodeOK().WithPayload(&models.Link{
				URL: &exampleLink,
			}),
		},
	}

	for _, tc := range testCases {
		links := &iservices.MockLinksService{}
		links.On("GetByCode", mock.MatchedBy(func(_ context.Context) bool { return true }), tc.params.Code).
			Return(tc.returnedLink, tc.returnedErr)

		h := NewGetLinksHandler(links)

		res := h.Handle(tc.params)

		assert.Equal(t, tc.res, res)
	}
}
