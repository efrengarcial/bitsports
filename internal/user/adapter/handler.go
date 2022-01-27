package adapter

import (
	"bitsports/internal/user"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler  represent the httphandler for user
type Handler struct {
	uc  *user.UseCase
}

// NewHandler will initialize the user/ resources endpoint
func NewHandler(uc *user.UseCase) *Handler {
	return &Handler{ uc }
}

//CreateUser create a new user
func (h *Handler) CreateUser(c echo.Context) error {
	var newUser user.NewUser

	if err := c.Bind(&newUser); err!=nil {
		return err
	}

	if err := c.Validate(newUser); err != nil {
		return err
	}

	u, err := h.uc.CreateUser(c.Request().Context(), newUser)
	if err!=nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) SignIn(c echo.Context) error {
	var token user.Token

	email, pass, ok := c.Request().BasicAuth()
	if !ok {
		err := errors.New("must provide email and password in Basic auth")
		return c.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized,"Unauthorized", err))
	}

	err := h.uc.Auth(c.Request().Context(), email, pass, &token)
	if err != nil {
		switch err {
		case user.ErrAuthenticationFailure:
			return c.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized,"Unauthorized"))
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, token)
}

