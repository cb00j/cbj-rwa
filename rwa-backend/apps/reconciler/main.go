package main

import (
	"context"
	"os"
	"time"

	config "github.com/cb00j/cbj-rwa/rwa-backend/apps/reconciler/config"
	"github.com/cb00j/cbj-rwa/rwa-backend/apps/reconciler/service"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/evm_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/database"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := &cli.Command{
		Name:        "reconciler",
		Description: "Retries stuck on-chain settlements and failed Kafka/WS event processing",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Value:   "./config/config.yaml",
				Aliases: []string{"c"},
				Usage:   "path to the config file",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			mainCtx := context.WithValue(ctx, log.TraceID, "reconciler")
			return startApp(mainCtx, c.String("config"))
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.ErrorZ(context.Background(), "app run failed", zap.Any("error", err))
	}
}

func startApp(ctx context.Context, configPath string) error {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		log.ErrorZ(ctx, "failed to load config file", zap.String("path", configPath), zap.Error(err))
		return err
	}
	log.InitLogger(conf.Logger)
	log.InfoZ(ctx, "load config file success", zap.String("path", configPath))

	cusLog := fx.WithLogger(func() fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log.Logger()}
	})

	app := fx.New(
		cusLog,
		config.LoadModule(conf),
		database.LoadModule(conf.Db),
		evm_helper.LoadModule(conf.RpcInfo),
		service.LoadModule(),
		fx.StartTimeout(120*time.Second),
		fx.StopTimeout(120*time.Second),
	)
	log.InfoZ(ctx, "Begin to start reconciler service")
	app.Run()
	return nil
}
