package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/practice/backend/intertnal/config"
	"github.com/practice/backend/intertnal/controller"
	"github.com/practice/backend/intertnal/controller/http"
	"github.com/practice/backend/intertnal/repository"
	"github.com/practice/backend/intertnal/repository/pgx"
	"github.com/practice/backend/pkg/postgres"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
)

func main() {
	var (
		err    error
		ctx    context.Context
		cfg    *config.Config
		log    *zap.Logger
		repo   repository.Interface
		server controller.Controller
		pool   *pgxpool.Pool
		pgErr  *pgconn.PgError
		cancel context.CancelFunc
	)

	ctx, cancel = signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	defer cancel()

	//initialize logger
	log, err = zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize logger", zap.Error(err))
	}

	//initialize config
	cfg, err = config.New(ctx)
	if err != nil {
		log.Fatal("Failed to initialize config", zap.Error(err))
	}

	//initialize pool
	pool, err = postgres.New(cfg, log)
	//initialize repository
	repo, err = pgx.New(pool, log, pgErr)
	if err != nil {
		log.Fatal("Failed to initialize database", zap.Error(err))
	}

	//initialize server
	server, err = http.New(log, cfg, repo)
	if err != nil {
		log.Fatal("Failed to initialize server", zap.Error(err))
	}

	//closing server
	defer func() {
		log.Error(
			"Shutting down server",
			zap.Error(server.Shutdown(ctx)),
		)
	}()

	err = server.Run(ctx)
	if err != nil {
		log.Fatal("Failed to run server", zap.Error(err))
	}

	<-ctx.Done()
	log.Info("[Graceful shutdown]")
}
