package pocketadapter

import (
	"context"
	"fmt"
	"net/url"
)

func NewMockAdapter() Adapter {
	return &mockAdapter{}
}

type mockAdapter struct{}

func (m *mockAdapter) AddPocketItem(ctx context.Context, url *url.URL) error {
	fmt.Println("URL", url.String())
	return nil
}
