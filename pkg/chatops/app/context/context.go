package context

import (
	"context"

	"github.com/UsingCoding/fpgo/pkg/maybe"

	"telegrambot/pkg/chatops/app/token"
)

type Context = context.Context

type ChatContext interface {
	Context

	ChatID() token.ChatID
	UserID() maybe.Maybe[token.UserID]
}

func NewChatContext(
	ctx Context,
	chatID token.ChatID,
	userID maybe.Maybe[token.UserID],
) ChatContext {
	return &chatContext{
		Context: ctx,
		chatID:  chatID,
		userID:  userID,
	}
}

type chatContext struct {
	Context

	chatID token.ChatID
	userID maybe.Maybe[token.UserID]
}

func (c *chatContext) ChatID() token.ChatID {
	return c.chatID
}

func (c *chatContext) UserID() maybe.Maybe[token.UserID] {
	return c.userID
}
