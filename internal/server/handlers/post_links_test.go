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

func TestNewPostLinksHandler(t *testing.T) {
	links := &iservices.MockLinksService{}
	validator := &iservices.MockURLValidatorService{}

	h := NewPostLinksHandler(links, validator)

	assert.Equal(t, links, h.links)
	assert.Equal(t, validator, h.validator)
}

func TestPostLinksHandler_Handle(t *testing.T) {
	codeExample := "example"
	linkExample := "https://example.com/"

	testCases := []struct {
		params operations.PostLinksParams

		returnedValidation bool
		returnedCode       string
		returnedErr        error

		res middleware.Responder
	}{
		{
			params: operations.PostLinksParams{Link: &models.Link{URL: &linkExample}},

			returnedValidation: false,

			res: operations.NewPostLinksBadRequest(),
		},
		{
			params: operations.PostLinksParams{Link: &models.Link{URL: &linkExample}},

			returnedValidation: true,
			returnedErr:        errors.New("unknown error"),

			res: operations.NewPostLinksInternalServerError(),
		},
		{
			params: operations.PostLinksParams{Link: &models.Link{URL: &linkExample}},

			returnedValidation: true,
			returnedCode:       codeExample,

			res: operations.NewPostLinksOK().WithPayload(&models.Code{Code: &codeExample}),
		},
	}

	for _, tc := range testCases {
		links := &iservices.MockLinksService{}
		links.On("Create", mock.MatchedBy(func(_ context.Context) bool { return true }), *tc.params.Link.URL).
			Return(tc.returnedCode, tc.returnedErr)

		validator := &iservices.MockURLValidatorService{}
		validator.On("Validate", *tc.params.Link.URL).Return(tc.returnedValidation)

		h := NewPostLinksHandler(links, validator)

		res := h.Handle(tc.params)

		assert.Equal(t, tc.res, res)
	}
}
