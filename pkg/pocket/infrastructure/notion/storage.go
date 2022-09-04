package notion

import (
	"context"

	"github.com/UsingCoding/fpgo/pkg/maybe"
	"github.com/jomei/notionapi"

	"telegrambot/pkg/pocket/app"
)

func NewStorage(secretKey, databaseID string) app.PocketItemStorage {
	return &storage{
		client:     notionapi.NewClient(notionapi.Token(secretKey)),
		databaseID: databaseID,
	}
}

type storage struct {
	client     *notionapi.Client
	databaseID string
}

func (s *storage) Store(ctx context.Context, item app.PocketItem) error {
	var imageURLStr string
	if maybe.Valid(item.ImageURL) {
		imageURLStr = maybe.Just(item.ImageURL).String()
	}

	// Create new page in Notion
	_, err := s.client.Page.Create(ctx, &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: notionapi.DatabaseID(s.databaseID),
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
			"URL": notionapi.URLProperty{
				Type: notionapi.PropertyTypeURL,
				URL:  item.URL.String(),
			},
		},
		Cover: &notionapi.Image{
			Type: notionapi.FileTypeExternal,
			External: &notionapi.FileObject{
				URL: imageURLStr,
			},
		},
	})
	return err
}
