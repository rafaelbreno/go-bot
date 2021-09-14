package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/config"
	"github.com/rafaelbreno/go-bot/api/handler"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/middlewares"
)

// Server manages HTTP
type Server struct {
	HTTP *fiber.App
	Ctx  *internal.Context
}

// Listen starts app
func NewServer() *Server {
	srv := &Server{
		HTTP: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: false,
			Concurrency:   256 * 1024,
			WriteTimeout:  time.Duration(45 * time.Second),
		}),
		Ctx: config.Ctx,
	}

	srv.routes()

	return srv
}

// ListenAndServe starts the web server.
func (s *Server) ListenAndServe() {
	if err := s.HTTP.Listen(s.Ctx.Env["API_PORT"]); err != nil {
		s.Ctx.Logger.Error(err.Error())
	}

}

// Close gracefully terminate the app.
func (s *Server) Close() {
	s.Ctx.Logger.Info("Gracefully terminating API...")
	if err := s.HTTP.Shutdown(); err != nil {
		s.Ctx.Logger.Error(err.Error())
	}
}

func (s *Server) routes() {
	s.commandRoutes()

	s.HTTP.Get("/ping", func(c *fiber.Ctx) error {
		s.Ctx.Logger.Info("GET /test")
		return c.JSON(map[string]string{
			"message": "PONG",
		})
	})

	a := s.HTTP.Group("/test", middlewares.CheckAuth)

	a.Get("/ping", func(c *fiber.Ctx) error {
		s.Ctx.Logger.Info("GET /test")
		s.Ctx.Logger.Info(c.UserContext().Value("user_id").(string))
		return c.JSON(map[string]string{
			"message": "PONG",
		})
	})
}

func (s *Server) commandRoutes() {
	commandGroup := s.HTTP.Group("/command", middlewares.CheckAuth)

	ch := handler.NewCommandHandler()

	commandGroup.Post("/create", ch.Create)
	commandGroup.Get("/:id", ch.Read)
	commandGroup.Put("/:id", ch.Update)
	commandGroup.Patch("/:id", ch.Update)
	commandGroup.Delete("/:id", ch.Delete)
}
