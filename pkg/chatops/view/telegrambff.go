package view

import (
	"time"

	"github.com/UsingCoding/fpgo/pkg/maybe"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"telegrambot/pkg/chatops/app"
	"telegrambot/pkg/chatops/app/auth"
	"telegrambot/pkg/chatops/app/context"
	"telegrambot/pkg/chatops/app/token"
	"telegrambot/pkg/common/infrastructure/logger"
)

type TelegramBFF interface {
	HandleUpdate(context context.Context, update tgbotapi.Update) tgbotapi.Chattable
}

func NewTelegramBFF(chatService app.ChatService, l logger.Logger) TelegramBFF {
	return &bff{chatService: chatService, logger: l}
}

type bff struct {
	chatService app.ChatService
	logger      logger.Logger
}

func (b *bff) HandleUpdate(ctx context.Context, update tgbotapi.Update) tgbotapi.Chattable {
	// For now just ignoring non message updates
	if update.Message == nil {
		return nil
	}

	userID := getUserID(update)

	chatCtx := context.NewChatContext(
		ctx,
		token.ChatID(update.Message.Chat.ID),
		userID,
	)

	err := b.logExecution("message", func() error {
		return b.chatService.HandleMessage(chatCtx, update.Message.Text)
	})

	if err != nil {
		return translateError(chatCtx, err)
	}

	return tgbotapi.NewMessage(int64(chatCtx.ChatID()), "Link added 👍")
}

func translateError(ctx context.ChatContext, err error) tgbotapi.Chattable {
	switch errors.Cause(err) {
	case app.ErrUnknownHandleAction:
		return tgbotapi.NewMessage(int64(ctx.ChatID()), "Unknown action to perform")
	case auth.ErrUnauthorizedUser:
		return tgbotapi.NewMessage(int64(ctx.ChatID()), "Unauthorized")
	default:
		return tgbotapi.NewMessage(int64(ctx.ChatID()), "Unknown error")
	}
}

func getUserID(update tgbotapi.Update) maybe.Maybe[token.UserID] {
	user := update.SentFrom()
	if user == nil {
		return maybe.NewNone[token.UserID]()
	}
	return maybe.NewJust[token.UserID](token.UserID(user.ID))
}

func (b *bff) logExecution(requestType string, f func() error) error {
	start := time.Now()
	err := f()

	fields := logger.Fields{
		"duration":    time.Since(start).String(),
		"requestType": requestType,
	}

	entry := b.logger.WithFields(fields)
	if err != nil {
		entry.Error(err, "call failed")
	} else {
		entry.Info("call finished")
	}

	return err
}
