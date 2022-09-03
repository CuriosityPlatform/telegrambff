package view

import (
	"github.com/UsingCoding/fpgo/pkg/maybe"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"telegrambot/pkg/chatops/app"
	"telegrambot/pkg/chatops/app/auth"
	"telegrambot/pkg/chatops/app/context"
	"telegrambot/pkg/chatops/app/token"
)

type TelegramBFF interface {
	HandleUpdate(context context.Context, update tgbotapi.Update) tgbotapi.Chattable
}

func NewTelegramBFF(chatService app.ChatService) TelegramBFF {
	return &bff{chatService: chatService}
}

type bff struct {
	chatService app.ChatService
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

	err := b.chatService.HandleMessage(chatCtx, update.Message.Text)
	if err != nil {
		return translateError(chatCtx, err)
	}

	return tgbotapi.NewMessage(int64(chatCtx.ChatID()), "Link added")
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
