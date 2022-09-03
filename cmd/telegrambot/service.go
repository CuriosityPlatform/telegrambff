package main

import (
	"context"

	"github.com/urfave/cli/v2"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"telegrambot/pkg/chatops"
	"telegrambot/pkg/chatops/app/token"
	"telegrambot/pkg/chatops/infrastructure/pocketadapter"
	"telegrambot/pkg/chatops/infrastructure/telegramserver"
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

	telegramBFF := chatops.Container(pocketadapter.NewMockAdapter(), token.UserID(config.AuthorizedUserID))

	updateListener := telegramserver.NewPullUpdateListener(bot)
	server := telegramserver.NewServer(
		updateListener,
		telegramBFF.HandleUpdate,
		nil,
	)

	return server.ListenAndServe(ctx)
}
