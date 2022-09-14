package chatops

import (
	"telegrambot/pkg/chatops/app"
	"telegrambot/pkg/chatops/app/auth"
	"telegrambot/pkg/chatops/app/token"
	"telegrambot/pkg/chatops/infrastructure/pocketadapter"
	"telegrambot/pkg/chatops/view"
	"telegrambot/pkg/common/infrastructure/logger"
)

func Container(pocket pocketadapter.Adapter, authorizedUserID token.UserID, l logger.Logger) view.TelegramBFF {
	chatService := app.NewChatService(pocket)
	chatService = auth.NewAuthService(authorizedUserID, chatService)
	return view.NewTelegramBFF(chatService, l)
}
