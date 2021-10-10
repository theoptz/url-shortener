package irepositories

import (
	"context"

	"github.com/theoptz/url-shortener/internal/repositories"
)

var _ = &MockRangeRepository{}

//go:generate mockery --all --inpackage --case underscore
type RangeRepository interface {
	Get(context.Context) (*repositories.RangeItem, error)
	GetNext(context.Context) (*repositories.RangeItem, error)
}
