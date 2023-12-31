package routes

import (
	"github.com/boltdbgui/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/fs"
	"net/http"
)

func RegisterStaticRoutes(e *echo.Echo) {
	staticPages := []string{"css", "styles", "img", "js", "app", "maps", "ico", "fonts", "video", "icons"}

	sub, err := fs.Sub(ui.WebContent, "dist")
	if err != nil {
		return
	}

	ac := http.FS(sub)

	var contentHandler = echo.WrapHandler(http.FileServer(ac))
	// set the root route (serving index.html)
	e.GET("/", contentHandler)

	// configure static routes
	for _, path := range staticPages {
		e.GET("/"+path+"/*", contentHandler, middleware.Gzip())
	}
}
