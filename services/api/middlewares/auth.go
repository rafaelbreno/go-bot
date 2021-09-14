package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-bot/api/config"
)

type userAPI struct {
	ID string `json:"id"`
}

// CheckAuth check if user is authenticated
func CheckAuth(ctx *fiber.Ctx) error {
	appCtx := config.Ctx

	b := ctx.Get("Authorization")
	bearer := string(b)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%s%s", appCtx.Env["AUTH_URL"], appCtx.Env["AUTH_PORT"], appCtx.Env["AUTH_ENDPOINT"]), nil)

	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	client := http.Client{}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	if res.StatusCode != http.StatusOK {
		return ctx.
			Status(res.StatusCode).
			JSON(fiber.Map{
				"error": "unauthorized",
			})
	}

	resBody := res.Body

	resBytes, _ := ioutil.ReadAll(resBody)

	var v userAPI

	if err := json.Unmarshal(resBytes, &v); err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	userCtx := context.WithValue(ctx.UserContext(), "user_id", v.ID)

	ctx.SetUserContext(userCtx)

	return ctx.Next()
}
