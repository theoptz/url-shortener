package links

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theoptz/url-shortener/internal/interfaces/irepositories"
	"github.com/theoptz/url-shortener/internal/interfaces/iservices"
	"github.com/theoptz/url-shortener/internal/repositories"
)

func TestLinks_GetByCode(t *testing.T) {
	testCases := []struct {
		code string

		returnedId   int64
		returnedLink string
		returnedErr  error

		expectedLink string
		expectedErr  error
	}{
		{
			returnedId: -1,
		},
		{
			returnedId:   -1,
			returnedLink: "http://example.com/",
			returnedErr:  errors.New("invalid id"),
		},
		{
			returnedId:  10,
			returnedErr: errors.New("not found"),

			expectedLink: "",
			expectedErr:  errors.New("not found"),
		},
		{
			returnedId:   10,
			returnedLink: "https://example.com/",

			expectedLink: "https://example.com/",
		},
	}

	for _, tc := range testCases {
		ctx := context.TODO()

		repo := &irepositories.MockLinksRepository{}
		counter := &iservices.MockCounterService{}
		coder := &iservices.MockCodeService{}

		coder.On("Decode", tc.code).Return(tc.returnedId)
		repo.On("GetByID", ctx, tc.returnedId).Return(tc.returnedLink, tc.returnedErr)

		links := NewLinks(repo, counter, coder)

		link, err := links.GetByCode(ctx, tc.code)

		assert.Equal(t, tc.expectedErr, err)
		assert.Equal(t, tc.expectedLink, link)
	}
}

func TestLinks_Create(t *testing.T) {
	testCases := []struct {
		link string

		returnedId   int64
		returnedCode string
		returnedErr  error

		expectedCode string
		expectedErr  error
	}{
		{
			returnedId:   1,
			returnedCode: "example",

			expectedCode: "example",
		},
		{
			returnedId:   1,
			returnedCode: "example",
			returnedErr:  errors.New("unknown error"),

			expectedCode: "",
			expectedErr:  errors.New("unknown error"),
		},
	}

	for _, tc := range testCases {
		ctx := context.TODO()

		repo := &irepositories.MockLinksRepository{}
		counter := &iservices.MockCounterService{}
		coder := &iservices.MockCodeService{}

		counter.On("Inc").Return(tc.returnedId, nil)
		coder.On("Encode", tc.returnedId).Return(tc.returnedCode)
		repo.On("Create", ctx, repositories.LinkItem{ID: tc.returnedId, Link: tc.link}).Return(tc.returnedErr)

		links := NewLinks(repo, counter, coder)

		code, err := links.Create(ctx, tc.link)

		assert.Equal(t, tc.expectedErr, err)
		assert.Equal(t, tc.expectedCode, code)
	}
}

func TestNewLinks(t *testing.T) {
	repo := &irepositories.MockLinksRepository{}
	counter := &iservices.MockCounterService{}
	coder := &iservices.MockCodeService{}

	res := NewLinks(repo, counter, coder)

	assert.Equal(t, repo, res.repo)
	assert.Equal(t, counter, res.counter)
	assert.Equal(t, coder, res.coder)
}
