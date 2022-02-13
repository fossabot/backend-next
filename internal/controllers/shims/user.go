package shims

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"github.com/penguin-statistics/backend-next/internal/models/shims"
	"github.com/penguin-statistics/backend-next/internal/server"
	"github.com/penguin-statistics/backend-next/internal/service"
	"github.com/penguin-statistics/backend-next/internal/utils"
)

type AccountController struct {
	service *service.AccountService
}

func RegisterAccountController(v2 *server.V2, s *service.AccountService) {
	c := &AccountController{
		service: s,
	}
	v2.Post("/users", c.Login)
}

// @Summary      Login with PenguinID
// @Tags         Account
// @Accept       plain
// @Produce      plain
// @Param        userId  body      int  true  "User ID"
// @Success      200     {array}  shims.SiteStats
// @Failure      500     {object}  errors.PenguinError "An unexpected error occurred"
// @Router       /PenguinStats/api/v2/users [POST]
// @Deprecated
func (c *AccountController) Login(ctx *fiber.Ctx) error {
	inputPenguinId := ctx.Body()

	account, err := c.service.GetAccountByPenguinId(ctx, string(inputPenguinId))
	if err != nil {
		return err
	}

	utils.SetPenguinIDToResponse(ctx, account.PenguinID)

	// for some reasons the response for the login API is in format of
	// text/plain so I'd have to manually convert it to JSON and use ctx#Send to respond
	// to ensure compatibility
	resp, err := json.Marshal(&shims.LoginResponse{
		UserID: account.PenguinID,
	})
	if err != nil {
		return err
	}

	utils.SetCtxNoCache(ctx)

	return ctx.Send(resp)
}