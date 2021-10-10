package irepositories

import (
	"context"

	"github.com/theoptz/url-shortener/internal/repositories"
)

var _ CounterRepository = &MockCounterRepository{}

// don't need "Update" method because counter increments automatically on inserting new link

//go:generate mockery --all --inpackage --case underscore
type CounterRepository interface {
	Get(context.Context, *repositories.RangeItem) (int64, error)
}
