package app

import (
	"context"
	"net/url"
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

func (s *service) AddPocketItem(ctx context.Context, url *url.URL) error {
	meta, err := s.metadataParser.Parse(url)
	if err != nil {
		return err
	}

	return s.pocketItemStorage.Store(ctx, PocketItem{
		Title:    meta.Title,
		URL:      url,
		ImageURL: meta.ImageURL,
	})
}
