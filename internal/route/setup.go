package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error { return c.String(http.StatusOK, "ok") }) //nolint:wrapcheck
}
