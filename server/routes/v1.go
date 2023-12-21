package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("", sayHello, can("api"))
}

func sayHello(c echo.Context) error {
	return nil
}

// can checks that the current user's role is allowed to perform all of the
// provided actions (so this is an AND condition, use canOr for OR)
func can(actions ...string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
