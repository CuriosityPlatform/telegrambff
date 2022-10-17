package app

import (
	"context"
	"net/url"

	"github.com/UsingCoding/fpgo/pkg/maybe"
)

type Service interface {
	AddPocketItem(ctx context.Context, url *url.URL) error
}

func NewService(pocketItemStorage PocketItemStorage, metadataParser MetadataParser) Service {
	return &service{pocketItemStorage: pocketItemStorage, metadataParser: metadataParser}
}

type service struct {
	pocketItemStorage PocketItemStorage
	metadataParser    MetadataParser
}

func (s *service) AddPocketItem(ctx context.Context, u *url.URL) error {
	meta, err := s.metadataParser.Parse(u)
	if err != nil {
		return err
	}

	// Take host as title if title not found
	title := u.Host
	if maybe.Valid(meta.Title) {
		title = maybe.Just(meta.Title)
	}

	return s.pocketItemStorage.Store(ctx, PocketItem{
		Title:    title,
		URL:      u,
		ImageURL: meta.ImageURL,
	})
}
