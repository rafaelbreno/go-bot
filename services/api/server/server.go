package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/handler"
	"github.com/rafaelbreno/go-bot/api/internal"
)

type Server struct {
	App  *fiber.App
	Ctx  *internal.Context
	Port string
}

func (s *Server) ListenAndServe(h handler.Handler) {
	s.App = fiber.New()

	if s.Port == "" {
		s.Ctx.Logger.Error("Empty port")
		os.Exit(0)
	}

	s.Port = string(append([]byte(":"), []byte(s.Port)...))

	s.App.Get("/test", func(c *fiber.Ctx) error {
		s.Ctx.Logger.Info("GET /test")
		return c.JSON(map[string]string{
			"message": "ok",
		})
	})

	s.App.Get("/ping", h.Ping)

	s.Ctx.Logger.Info(fmt.Sprintf("Starting server, listening port: %s", s.Port))
	if err := s.App.Listen(s.Port); err != nil {
		s.Ctx.Logger.Error(err.Error())
	}
}

func (s *Server) Close() {
	s.Ctx.Logger.Info("Shutting down server...")
	if err := s.App.Shutdown(); err != nil {
		s.Ctx.Logger.Error(err.Error())
	}
	s.Ctx.Logger.Info("Server shut down!")
}
