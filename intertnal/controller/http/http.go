package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Controller struct {
	server *echo.Echo
	log    *zap.Logger
	cfg    config.Config
	repo   repository.Interface
}

func New(
	log *zap.Logger,
	cfg config.Config,
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

func (ctrl *Controller) configure() {

}
