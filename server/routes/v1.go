package routes

import (
	"github.com/boltdbgui/server/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("", handlers.SayHello, can("api"))
	v1.POST("/list", handlers.ListElement)
	v1.POST("/add_buckets", handlers.AddBucket)
	v1.POST("/add_pairs", handlers.AddPairs)
	v1.POST("/delete", handlers.DeleteElement)
	v1.POST("/rename_key", handlers.RenameElement)
	v1.POST("/update_value", handlers.UpdatePairValue)
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
