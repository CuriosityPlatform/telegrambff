package main

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/urfave/cli/v2"

	"telegrambot/pkg/chatops"
	"telegrambot/pkg/chatops/app/token"
	"telegrambot/pkg/chatops/infrastructure/telegramserver"
	"telegrambot/pkg/pocket"
)

func service(config *config) *cli.Command {
	return &cli.Command{
		Name: "service",
		Action: func(c *cli.Context) error {
			return runService(c.Context, config)
		},
	}
}

func runService(ctx context.Context, config *config) error {
	bot, err := tgbotapi.NewBotAPI(config.TelegramAPIToken)
	if err != nil {
		return err
	}

	pocketAPI := pocket.ContainerAPI(config.NotionSecretToken, config.NotionDatabaseID)
	telegramBFF := chatops.Container(pocketAPI, token.UserID(config.AuthorizedUserID))

	updateListener := telegramserver.NewPullUpdateListener(bot)
	server := telegramserver.NewServer(
		updateListener,
		telegramBFF.HandleUpdate,
		nil,
	)

	return server.ListenAndServe(ctx)
}
