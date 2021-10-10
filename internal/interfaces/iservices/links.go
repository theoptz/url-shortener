package iservices

import "context"

var _ LinksService = &MockLinksService{}

//go:generate mockery --all --inpackage --case underscore
type LinksService interface {
	GetByCode(context.Context, string) (string, error)
	Create(context.Context, string) (string, error)
}
