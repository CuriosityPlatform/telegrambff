package notion

import (
	"context"

	"github.com/UsingCoding/fpgo/pkg/maybe"
	"github.com/jomei/notionapi"
	"github.com/pkg/errors"

	"telegrambot/pkg/pocket/app"
)

const (
	webResourceType = "Web resource"
)

func NewWebResourceStorage(secretKey, databaseID string) app.PocketItemStorage {
	return &webResourceStorage{
		client:     notionapi.NewClient(notionapi.Token(secretKey)),
		databaseID: databaseID,
	}
}

type webResourceStorage struct {
	client     *notionapi.Client
	databaseID string
}

func (storage *webResourceStorage) Store(ctx context.Context, item app.PocketItem) error {
	var coverImage *notionapi.Image

	if maybe.Valid(item.ImageURL) {
		coverImage = &notionapi.Image{
			Type: notionapi.FileTypeExternal,
			External: &notionapi.FileObject{
				URL: maybe.Just(item.ImageURL).String(),
			},
		}
	}

	// Create new page in Notion
	_, err := storage.client.Page.Create(ctx, &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: notionapi.DatabaseID(storage.databaseID),
		},
		Properties: map[string]notionapi.Property{
			"Name": notionapi.TitleProperty{
				Type: notionapi.PropertyTypeTitle,
				Title: []notionapi.RichText{{
					Type: notionapi.ObjectTypeText,
					Text: notionapi.Text{
						Content: item.Title,
					},
				}},
			},
			"Resource origin": notionapi.SelectProperty{
				Type: notionapi.PropertyTypeSelect,
				Select: notionapi.Option{
					Name: webResourceType,
				},
			},
			"URL": notionapi.URLProperty{
				Type: notionapi.PropertyTypeURL,
				URL:  item.URL.String(),
			},
		},
		Cover: coverImage,
	})
	return errors.WithStack(err)
}
