package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/entity"
	"github.com/rafaelbreno/go-bot/api/repository"
)

// CommandHandler manages all endpoints
// related to Command entity.
type CommandHandler struct {
	repo repository.CommandRepoCtx
}

// NewCommandHandler created and returns
// a configured CommandHandler.
func NewCommandHandler() CommandHandler {
	return CommandHandler{repository.NewCommandRepoCtx()}
}

// Create - receive and creates a new command
func (h *CommandHandler) Create(c *fiber.Ctx) error {
	commandJSON := new(entity.CommandJSON)

	if len(c.Body()) <= 0 {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": "Empty body",
			})
	}

	if err := c.BodyParser(commandJSON); err != nil {
		h.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	command, err := h.repo.Create(commandJSON.ToCommand())

	if err != nil {
		h.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.
		Status(http.StatusCreated).
		JSON(command)
}

// Read - return a Command with given ID
func (h *CommandHandler) Read(c *fiber.Ctx) error {
	command, err := h.repo.Read(c.Params("id"))

	if err != nil {
		h.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusNotFound).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.
		Status(http.StatusOK).
		JSON(command)
}

// Update - receive Command's fields and ID
// to update it
func (h *CommandHandler) Update(c *fiber.Ctx) error {
	if len(c.Body()) <= 0 {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": "Empty body",
			})
	}
	commandJSON := new(entity.CommandJSON)

	if err := c.BodyParser(commandJSON); err != nil {
		h.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	command, err := h.repo.Update(c.Params("id"), commandJSON.ToCommand())

	if err != nil {
		h.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.
		Status(http.StatusOK).
		JSON(command)
}

// Delete - receive a id and delete the
// command identified by it
func (h *CommandHandler) Delete(c *fiber.Ctx) error {
	err := h.repo.Delete(c.Params("id"))

	if err != nil {
		h.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusNotFound).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"message": "Command deleted!",
		})
}
