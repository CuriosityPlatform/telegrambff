package chatops

import (
	"telegrambot/pkg/chatops/app"
	"telegrambot/pkg/chatops/app/auth"
	"telegrambot/pkg/chatops/app/token"
	"telegrambot/pkg/chatops/infrastructure/pocketadapter"
	"telegrambot/pkg/chatops/view"
)

func Container(pocket pocketadapter.Adapter, authorizedUserID token.UserID) view.TelegramBFF {
	chatService := app.NewChatService(pocket)
	chatService = auth.NewAuthService(authorizedUserID, chatService)
	return view.NewTelegramBFF(chatService)
}
