package pocket

import (
	"telegrambot/pkg/pocket/api"
	"telegrambot/pkg/pocket/app"
	"telegrambot/pkg/pocket/infrastructure/metadataparser"
	"telegrambot/pkg/pocket/infrastructure/notion"
)

func ContainerAPI(notionSecretToken, notionDatabaseID, notionResourceDatabaseID string) api.API {
	return api.NewAPI(app.NewService(
		notion.NewCompositeStorage(
			notion.NewWebResourceStorage(notionSecretToken, notionResourceDatabaseID),
			notion.NewStorage(notionSecretToken, notionDatabaseID),
		),
		metadataparser.NewOpenGraphParser(),
	))
}
