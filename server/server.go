package server

import (
	"net/http"

	"github.com/boltdbgui/server/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// StartServer initializes and starts the HTTP server on the specified port.
func StartServer(port string) {
	// Initialize Echo instance
	e := echo.New()

	// Apply middleware
	e.Use(middleware.CORS())
	// Uncomment and apply middleware as needed
	// e.Use(middleware.Recover())

	// Register routes
	registerRoutes(e)

	// Configure server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: e,
	}

	// Start server and handle errors
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		e.Logger.Fatalf("Failed to start server: %v", err)
	}
}

// registerRoutes registers all routes with the Echo instance.
func registerRoutes(e *echo.Echo) {
	routes.RegisterStaticRoutes(e)
	routes.RegisterV1Routes(e)
}
