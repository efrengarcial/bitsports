package errorhandling

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type iErrRecordNotFound interface {
	error
	GetMessageWithID() string
}

type iErrDuplicateRecord interface {
	error
	GetMessageDuplicate() string
}

type iErrBadRequest interface {
	error
	GetTitle() string
	GetDetails() interface{}
}

// Error Handler in server package
func Error(logger *logrus.Logger) func(error, echo.Context) {

	return func(err error, c echo.Context) {
		var report *echo.HTTPError
		if c.Response().Committed {
			return
		}

		switch err.(type) {
		case iErrRecordNotFound:
			er := err.(iErrRecordNotFound)
			report = echo.NewHTTPError(http.StatusNotFound, er.GetMessageWithID())
		case iErrDuplicateRecord:
			er := err.(iErrDuplicateRecord)
			report = echo.NewHTTPError(http.StatusConflict, er.GetMessageDuplicate())
		case iErrBadRequest:
			er := err.(iErrBadRequest)
			report = echo.NewHTTPError(http.StatusBadRequest, er.Error())
			logger.Debug(fmt.Sprintf("Error:%+v %+v", er, er.GetDetails()))
		case validator.ValidationErrors:
			if castedObject, ok := err.(validator.ValidationErrors); ok {
				message := make(map[string]string)
				for _, err := range castedObject {
					switch err.Tag() {
					case "alpha":
						message[strings.ToLower(err.Field())] = fmt.Sprintf("'%s' campo solo permite caracteres.",
							err.Field())
					case "required":
						message[strings.ToLower(err.Field())] = fmt.Sprintf("'%s' campo es requerido",
							err.Field())
					case "gte":
						message[strings.ToLower(err.Field())] = fmt.Sprintf("'%s' longitud del campo debe ser mayor o igual a %s",
							err.Field(), err.Param())
					case "lte":
						message[strings.ToLower(err.Field())] = fmt.Sprintf("'%s' longitud del campo debe ser menor o igual a %s",
							err.Field(), err.Param())
					case "eqfield":
						message[strings.ToLower(err.Field())] = fmt.Sprintf("'%s' se debe ingresar la confirmación del campo %s",
							err.Field(), err.Param())
					case "email":
						message[strings.ToLower(err.Field())] = fmt.Sprintf("'%s' no es un correo válido",
							err.Field())
					}

				}
				report = echo.NewHTTPError(http.StatusBadRequest, message)
			}
		default:
			var ok bool
			report, ok = err.(*echo.HTTPError)
			if ok {
				if report.Message == "missing or malformed jwt" {
					report = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
				} else {
					report = echo.NewHTTPError(http.StatusBadRequest, "Request invalida")
				}

				logger.Debug(fmt.Sprintf("Error: %+v", err))
			} else {
				report = echo.NewHTTPError(http.StatusInternalServerError, "Error Interno del Servidor")
				logger.Error(fmt.Sprintf("Error: %+v", err))
			}

		}

		c.JSON(report.Code, report)
	}
}
