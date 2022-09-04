package api

import (
	"context"
	"net/url"

	"telegrambot/pkg/pocket/app"
)

type API interface {
	AddPocketItem(ctx context.Context, url *url.URL) error
}

func NewAPI(s app.Service) API {
	return &api{s: s}
}

type api struct {
	s app.Service
}

func (a *api) AddPocketItem(ctx context.Context, u *url.URL) error {
	return a.s.AddPocketItem(ctx, u)
}
