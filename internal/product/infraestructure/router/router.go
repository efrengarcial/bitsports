package router

import (
	"bitsports/pkg/auth"
	"bitsports/pkg/errorhandling"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// _queryPath Path of route
const _queryPath = "/graphql"

// New creates route endpoint
func New(h *handler.Handler, logger *logrus.Logger) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = errorhandling.Error(logger)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Restricted from here
	r := e.Group(_queryPath)
	key, err := auth.GetRSAPublicKey()
	if err != nil {
		logger.Fatal(err)
	}

	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    key,
		SigningMethod: "RS256",
	}))

	r.POST("", echo.WrapHandler(h))

	return e
}
