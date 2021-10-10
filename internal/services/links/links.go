package links

import (
	"context"

	"github.com/theoptz/url-shortener/internal/interfaces/irepositories"
	"github.com/theoptz/url-shortener/internal/interfaces/iservices"
	"github.com/theoptz/url-shortener/internal/repositories"
)

type Links struct {
	repo    irepositories.LinksRepository
	counter iservices.CounterService
	coder   iservices.CodeService
}

func (l *Links) GetByCode(ctx context.Context, code string) (string, error) {
	id := l.coder.Decode(code)
	if id < 0 {
		return "", nil
	}

	return l.repo.GetByID(ctx, id)
}

func (l *Links) Create(ctx context.Context, link string) (string, error) {
	id, err := l.counter.Inc()
	if err != nil {
		return "", err
	}

	code := l.coder.Encode(id)

	if err := l.repo.Create(ctx, repositories.LinkItem{ID: id, Link: link}); err != nil {
		return "", err
	}

	return code, nil
}

func NewLinks(repo irepositories.LinksRepository, counter iservices.CounterService, coder iservices.CodeService) *Links {
	return &Links{
		repo:    repo,
		coder:   coder,
		counter: counter,
	}
}
