package infraestructure

import (
	"bitsports/internal/user/adapter"
	"bitsports/pkg/errorhandling"
	"bitsports/pkg/validate"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, h *adapter.Handler, logger *logrus.Logger) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//link our validator to echo framework
	e.Validator = &validate.Validator{}

	e.HTTPErrorHandler = errorhandling.Error(logger)

	e.GET("/token", func(context echo.Context) error { return h.SignIn(context) })
	e.POST("/users", func(context echo.Context) error { return h.CreateUser(context) })

	return e
}
