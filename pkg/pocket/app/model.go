package app

import (
	"context"
	"net/url"

	"github.com/UsingCoding/fpgo/pkg/maybe"
)

type PocketItem struct {
	Title    string
	URL      *url.URL
	ImageURL maybe.Maybe[*url.URL]
}

type PocketItemStorage interface {
	Store(ctx context.Context, item PocketItem) error
}
