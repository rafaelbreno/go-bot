package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/handler"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/storage"
)

type Server struct {
	App     *fiber.App
	Ctx     *internal.Context
	Port    string
	Handler handler.Handler
}

func (s *Server) ListenAndServe(h *handler.Handler) {
	s.App = fiber.New()

	st := storage.NewStorage(s.Ctx)

	h.Storage = st

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

	s.CommandServer()

	s.Ctx.Logger.Info(fmt.Sprintf("Starting server, listening port: %s", s.Port))
	if err := s.App.Listen(s.Port); err != nil {
		s.Ctx.Logger.Error(err.Error())
	}
}

func (s *Server) CommandServer() {
	uh := handler.CommandHandler{
		Ctx:     s.Handler.Ctx,
		Storage: storage.NewStorage(s.Ctx),
	}

	_, err := uh.Storage.SQL.Client.DB()
	if err != nil {
		s.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}

	u := s.App.Group("/command")

	u.Post("/create", uh.Create)
}

func (s *Server) Close() {
	s.Ctx.Logger.Info("Shutting down server...")
	if err := s.App.Shutdown(); err != nil {
		s.Ctx.Logger.Error(err.Error())
	}
	s.Ctx.Logger.Info("Server shut down!")
}
