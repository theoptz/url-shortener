package irepositories

import (
	"context"

	"github.com/theoptz/url-shortener/internal/repositories"
)

var _ LinksRepository = &MockLinksRepository{}

//go:generate mockery --all --inpackage --case underscore
type LinksRepository interface {
	Create(context.Context, repositories.LinkItem) error
	GetByID(context.Context, int64) (string, error)
}
