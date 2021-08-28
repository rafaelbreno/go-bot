package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/storage"
)

// Handler handles http requests
type Handler struct {
	Ctx     *internal.Context
	Storage *storage.Storage
}

type keyMap map[string]string

// Ping returns ok
func (h Handler) Ping(c *fiber.Ctx) error {
	h.Ctx.Logger.Info("GET /ping")
	return c.JSON(keyMap{
		"message": "ok",
	})
}
