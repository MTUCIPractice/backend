package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *Controller) Ping(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		fmt.Sprint("pong"),
	)
}
