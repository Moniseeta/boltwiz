package routes

import (
	"io/fs"
	"net/http"

	"github.com/boltdbgui/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RegisterStaticRoutes registers static routes for the Echo instance.
func RegisterStaticRoutes(e *echo.Echo) {
	staticPages := []string{"css", "styles", "img", "js", "app", "maps", "ico", "fonts", "video", "icons"}

	sub, err := fs.Sub(ui.WebContent, "dist")
	if err != nil {
		e.Logger.Fatalf("Failed to subtree web content: %v", err)
		return
	}

	ac := http.FS(sub)
	contentHandler := echo.WrapHandler(http.FileServer(ac))

	// Set the root route (serving index.html)
	e.GET("/", contentHandler)

	// Configure static routes
	for _, path := range staticPages {
		e.GET("/"+path+"/*", contentHandler, middleware.Gzip())
	}
}
