package server

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/boltdbgui/server/routes"

	"github.com/labstack/echo/v4"
)

func StartServer(port string) {
	// Echo instance
	e := echo.New()

	//e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// Routes
	routes.RegisterStaticRoutes(e)
	routes.RegisterV1Routes(e)

	server := &http.Server{
		Addr: ":" + port,
	}

	if err := e.StartServer(server); err != http.ErrServerClosed {
		e.Logger.Fatalf("Failed to start server, %v", err)
	}
}
