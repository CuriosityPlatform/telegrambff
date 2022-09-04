package pocket

import (
	"telegrambot/pkg/pocket/api"
	"telegrambot/pkg/pocket/app"
	"telegrambot/pkg/pocket/infrastructure/metadataparser"
	"telegrambot/pkg/pocket/infrastructure/notion"
)

func ContainerAPI(notionSecretToken, notionDatabaseID string) api.API {
	return api.NewAPI(app.NewService(
		notion.NewStorage(notionSecretToken, notionDatabaseID),
		metadataparser.NewOpenGraphParser(),
	))
}
