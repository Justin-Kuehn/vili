package auth

import (
	"net/http"

	"github.com/airware/vili/server"
	"github.com/airware/vili/session"
	"github.com/labstack/echo"
)

var service Service

// Service is an authentication service interface
type Service interface {
	AddHandlers(s *server.Server)
	Cleanup()
}

// AddHandlers adds auth handlers to the server
func AddHandlers(s *server.Server) {
	service.AddHandlers(s)
	s.Echo().GET("/logout", logoutHandler)
}

// Cleanup cleans up the auth service
func Cleanup() {
	service.Cleanup()
}

func logoutHandler(c echo.Context) error {
	err := session.Logout(c.Request(), c.Response())
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/")
}
