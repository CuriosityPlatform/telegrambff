package auth

import (
	"github.com/UsingCoding/fpgo/pkg/maybe"
	"github.com/pkg/errors"

	"telegrambot/pkg/chatops/app"
	"telegrambot/pkg/chatops/app/context"
	"telegrambot/pkg/chatops/app/token"
)

var (
	ErrUnauthorizedUser = errors.New("unauthorized user")
)

func NewAuthService(authorizedUserID token.UserID, service app.ChatService) app.ChatService {
	return &authService{authorizedUserID: authorizedUserID, service: service}
}

type authService struct {
	authorizedUserID token.UserID

	service app.ChatService
}

func (s *authService) HandleMessage(ctx context.ChatContext, text string) error {
	if !maybe.Valid(ctx.UserID()) {
		return ErrUnauthorizedUser
	}

	if maybe.Just(ctx.UserID()) != s.authorizedUserID {
		return ErrUnauthorizedUser
	}

	return s.service.HandleMessage(ctx, text)
}
