package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/entity"
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

// CommandHandler handlers http requests
// related to user
type CommandHandler struct {
	Ctx     *internal.Context
	Storage *storage.Storage
}

// Create a command
func (h *CommandHandler) Create(c *fiber.Ctx) error {
	commandJSON := new(entity.CommandJSON)

	if len(c.Body()) <= 0 {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(keyMap{
				"error": "Empty body",
			})
	}

	if err := c.BodyParser(commandJSON); err != nil {
		h.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(keyMap{
				"error": err.Error(),
			})
	}

	command := commandJSON.ToCommand()

	// arruma pufavo >.<

	if err := h.
		Storage.
		SQL.
		Client.
		Create(&command).
		Error; err != nil {
		h.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(keyMap{
				"error": err.Error(),
			})
	}

	return c.
		JSON(command.ToJSON())
}
