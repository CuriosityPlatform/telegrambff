package pocketadapter

import (
	"context"
	"net/url"

	"telegrambot/pkg/pocket/api"
)

type Adapter interface {
	AddPocketItem(ctx context.Context, url *url.URL) error
}

func NewAdapter(pocketAPI api.API) Adapter {
	return &adapter{pocketAPI: pocketAPI}
}

type adapter struct {
	pocketAPI api.API
}

func (a *adapter) AddPocketItem(ctx context.Context, u *url.URL) error {
	return a.pocketAPI.AddPocketItem(ctx, u)
}
