package notion

import (
	"context"

	"github.com/pkg/errors"

	"telegrambot/pkg/pocket/app"
)

func NewCompositeStorage(storages ...app.PocketItemStorage) app.PocketItemStorage {
	return &compositeStorage{storages: storages}
}

type compositeStorage struct {
	storages []app.PocketItemStorage
}

func (s *compositeStorage) Store(ctx context.Context, item app.PocketItem) error {
	for _, storage := range s.storages {
		err := storage.Store(ctx, item)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
