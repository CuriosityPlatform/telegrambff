package app

import (
	"net/url"

	"telegrambot/pkg/chatops/app/context"
	"telegrambot/pkg/chatops/infrastructure/pocketadapter"
)

type ChatService interface {
	HandleMessage(ctx context.ChatContext, text string) error
}

func NewChatService(pocketAdapter pocketadapter.Adapter) ChatService {
	return &chatService{pocketAdapter: pocketAdapter}
}

type chatService struct {
	pocketAdapter pocketadapter.Adapter
}

func (service *chatService) HandleMessage(ctx context.ChatContext, text string) error {
	if textURL, err := url.Parse(text); err == nil {
		// Just add url to pocket
		return service.pocketAdapter.AddPocketItem(ctx, textURL)
	}

	return ErrUnknownHandleAction
}
