package main

import (
	"context"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"

	jsonlogger "telegrambot/pkg/common/infrastructure/logger"
)

const (
	appID = "telegrambot"
)

var (
	version = "UNKNOWN"
)

func main() {
	ctx := context.Background()

	ctx = subscribeForKillSignals(ctx)

	err := runApp(ctx, append(os.Args, "service"))
	if err != nil {
		stdlog.Fatal(err)
	}
}

func runApp(ctx context.Context, args []string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger, err := initLogger()
	if err != nil {
		return err
	}

	c, err := parseEnv()
	if err != nil {
		return err
	}

	app := &cli.App{
		Name:    appID,
		Version: version,
		Commands: []*cli.Command{
			service(c, logger),
		},
	}

	return app.RunContext(ctx, args)
}

func subscribeForKillSignals(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer cancel()
		select {
		case <-ctx.Done():
			signal.Stop(ch)
		case <-ch:
		}
	}()

	return ctx
}
func initLogger() (jsonlogger.MainLogger, error) {
	return jsonlogger.NewLogger(&jsonlogger.Config{AppName: appID}), nil
}
