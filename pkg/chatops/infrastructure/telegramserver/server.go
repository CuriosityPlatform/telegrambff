package telegramserver

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Server interface {
	ListenAndServe(context.Context) error
}

func NewServer(
	listener UpdatesListener,
	handler func(ctx context.Context, update tgbotapi.Update) tgbotapi.Chattable,
	requestContextFunc func(ctx context.Context, update tgbotapi.Update) context.Context,
) Server {
	return &server{listener: listener, Handler: handler, RequestContextFunc: requestContextFunc}
}

type server struct {
	listener UpdatesListener

	Handler            func(ctx context.Context, update tgbotapi.Update) tgbotapi.Chattable
	RequestContextFunc func(ctx context.Context, update tgbotapi.Update) context.Context
}

func (s *server) ListenAndServe(ctx context.Context) error {
	return s.listener.Listen(ctx, s.serve)
}

func (s *server) serve(ctx context.Context, update tgbotapi.Update) tgbotapi.Chattable {
	if s.RequestContextFunc != nil {
		ctx = s.RequestContextFunc(ctx, update)
	}

	return s.Handler(ctx, update)
}
