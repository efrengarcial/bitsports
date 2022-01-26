package router

import (
	"bitsports/pkg/errorhandling"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// QueryPath Path of route
const (
	QueryPath      = "/graphql"
)

// New creates route endpoint
func New(srv *handler.Handler, logger *logrus.Logger) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = errorhandling.Error(logger)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.POST(QueryPath, echo.WrapHandler(srv))

	return e
}
