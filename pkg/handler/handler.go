package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// GraphQLHandler handle handler wrapper between go-graphql relay with echo
func GraphQLHandler(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {

		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
