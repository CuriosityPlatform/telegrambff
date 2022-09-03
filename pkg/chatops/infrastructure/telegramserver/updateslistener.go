package telegramserver

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdatesListener interface {
	Listen(ctx context.Context, f ListenFunc) error
}

type ListenFunc func(ctx context.Context, update tgbotapi.Update) tgbotapi.Chattable

func NewPullUpdateListener(botAPI *tgbotapi.BotAPI) *PullUpdateListener {
	return &PullUpdateListener{botAPI: botAPI}
}

type PullUpdateListener struct {
	botAPI *tgbotapi.BotAPI
	logger log.Logger
}

func (listener *PullUpdateListener) Listen(ctx context.Context, f ListenFunc) error {
	updates := listener.botAPI.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:         0,
		Limit:          0,
		Timeout:        0,
		AllowedUpdates: nil,
	})

	defer listener.botAPI.StopReceivingUpdates()

	for {
		select {
		case update := <-updates:
			go func() {
				response := f(ctx, update)
				_, err := listener.botAPI.Send(response)
				if err != nil {
					listener.logger.Println("err while sending response to telegram", err)
				}
			}()
		case <-ctx.Done():
			return nil
		}
	}
}
