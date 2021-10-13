package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/entity"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/repository"
)

type AuthHandler struct {
	repo *repository.AuthRepoCtx
}

func NewUserHandler(ctx *internal.Context) AuthHandler {
	return AuthHandler{repository.NewAuthRepoCtx(ctx)}
}

func (a *AuthHandler) Create(c *fiber.Ctx) error {
	userJSON := new(entity.User)

	if len(c.Body()) <= 0 {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": "Empty body",
			})
	}

	if err := c.BodyParser(userJSON); err != nil {
		a.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	authRes, err := a.repo.Create(*userJSON)

	if err != nil {
		a.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	return c.
		Status(http.StatusCreated).
		JSON(authRes)
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	userJSON := new(entity.User)

	if len(c.Body()) <= 0 {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": "Empty body",
			})
	}

	if err := c.BodyParser(userJSON); err != nil {
		a.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	authRes, err := a.repo.Login(*userJSON)

	if err != nil {
		a.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	return c.
		Status(http.StatusCreated).
		JSON(authRes)
}

func (a *AuthHandler) Check(c *fiber.Ctx) error {
	userJSON := new(entity.User)

	if len(c.Body()) <= 0 {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": "Empty body",
			})
	}

	if err := c.BodyParser(userJSON); err != nil {
		a.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	err := a.repo.Check(userJSON.Token)

	if err != nil {
		a.repo.Ctx.Logger.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	return c.
		Status(http.StatusCreated).
		JSON(fiber.Map{
			"message": "ok",
		})
}
