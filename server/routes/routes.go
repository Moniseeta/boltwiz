package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterStaticRoutes(e *echo.Echo) {
	staticPages := []string{"css", "styles", "img", "js", "app", "maps", "ico", "fonts", "video", "icons"}

	// set the root route (serving index.html)
	e.GET("/", func(c echo.Context) error {
		return c.File("./ui/dist/index.html")
	})

	// configure static routes
	for _, path := range staticPages {
		e.Group("/"+path, middleware.Gzip()).
			Use(
				func(next echo.HandlerFunc) echo.HandlerFunc {
					return func(c echo.Context) error {
						c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
						return next(c)
					}
				},
				middleware.StaticWithConfig(
					middleware.StaticConfig{
						Root: "./ui/dist/" + path,
					}),
			)
	}
}
