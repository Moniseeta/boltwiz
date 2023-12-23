package server

import (
	"net/http"

	"github.com/boltdbgui/server/routes"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	// Echo instance
	e := echo.New()

	//e.Use(middleware.Recover())
	// Routes
	routes.RegisterStaticRoutes(e)
	routes.RegisterV1Routes(e)

	server := &http.Server{
		Addr: ":8090",
	}

	if err := e.StartServer(server); err != http.ErrServerClosed {
		e.Logger.Fatalf("Failed to start server, %v", err)
	}
}
