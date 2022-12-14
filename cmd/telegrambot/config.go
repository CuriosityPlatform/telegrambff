package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func parseEnv() (*config, error) {
	c := new(config)
	if err := envconfig.Process(appID, c); err != nil {
		return nil, errors.Wrap(err, "failed to parse env")
	}
	return c, nil
}

type config struct {
	TelegramAPIToken         string `envconfig:"telegram_api_token"`
	AuthorizedUserID         int64  `envconfig:"authorized_user_id"`
	NotionSecretToken        string `envconfig:"notion_secret_token"`
	NotionDatabaseID         string `envconfig:"notion_database_id"`
	NotionResourceDatabaseID string `envconfig:"notion_resource_database_id"`
}
