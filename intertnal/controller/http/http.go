package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/practice/backend/intertnal/config"
	"github.com/practice/backend/intertnal/controller"
	"github.com/practice/backend/intertnal/repository"
	"go.uber.org/zap"
)

var _ controller.Controller = (*Controller)(nil)

type Controller struct {
	server *echo.Echo
	log    *zap.Logger
	cfg    *config.Config
	repo   repository.Interface
}

func New(
	log *zap.Logger,
	cfg *config.Config,
	repo repository.Interface,
) (*Controller, error) {
	log.Info("Creating http controller")
	ctrl := &Controller{
		server: echo.New(),
		log:    log,
		cfg:    cfg,
		repo:   repo,
	}

	if err := ctrl.configure(); err != nil {

		return nil, err
	}
	return ctrl, nil
}

func (ctrl *Controller) configure() error {
	ctrl.configureRoutes()
	//ctrl.configureMiddlewares()
	return nil
}

func (ctrl *Controller) configureRoutes() {
	log.Info("Configuring routes")
	api := ctrl.server.Group("/api")
	{
		api.GET("/ping", ctrl.Ping)
	}
}

func (ctrl *Controller) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		ctrl.log.Info("Starting HTTP server on address", zap.String("", ctrl.cfg.Controller.GetBindAddress()))
		err := ctrl.server.Start(ctrl.cfg.Controller.GetBindAddress())
		if err != nil {
			cancel()
		}
	}()
	return ctx.Err()
}

func (ctrl *Controller) Shutdown(ctx context.Context) error {
	return ctrl.server.Shutdown(ctx)
}
