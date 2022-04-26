package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Custom(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside custom middleware")
		next(c)
		return nil
	}
}
