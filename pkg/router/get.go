package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/afifurrohman-id/tempsy-client/internal/client/auth"
	"github.com/afifurrohman-id/tempsy-client/internal/client/models"
	"github.com/afifurrohman-id/tempsy-client/internal/client/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleWelcomeClient(ctx *fiber.Ctx) error {
	return ctx.Render("pages/index", map[string]any{
		"user": ctx.Locals("user"),
	})
}

func HandleDashboardClient(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)
	if !ok {
		log.Panic("invalid_user_struct")
	}

	agent := fiber.Get(fmt.Sprintf("%s/files/%s?limit=%d&name=%s&size=%d&mime_type=%s", os.Getenv("API_SERVER_URL"), ctx.Params("username"), ctx.QueryInt("limit"), ctx.Query("name"), ctx.QueryInt("size"), ctx.Query("mime_type")))

	agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+ctx.Locals("token").(string))

	apiRes := new([]*models.DataFile)

	statusCode, body, errs := agent.Struct(&apiRes)
	if len(errs) > 0 {
		utils.Check(errs[0])
	}

	if statusCode != fiber.StatusOK {
		apiErr := new(models.ApiError)
		utils.Check(json.Unmarshal(body, &apiErr))

		return ctx.Render("pages/error", map[string]any{
			"code":    statusCode,
			"message": apiErr.Description,
		})
	}

	return ctx.Render("pages/dashboard", map[string]any{
		"user":  user,
		"files": apiRes,
		"type":  "Upload",
	})
}

func HandleProfileClient(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)
	if !ok {
		log.Panic("invalid_user_struct")
	}

	lastLoginMs, err := strconv.ParseInt(ctx.Cookies("last_login", fmt.Sprintf("%d", time.Now().UnixMilli())), 10, 64)
	utils.Check(err)

	apiRes, err := auth.GetUserInfo(ctx.Locals("token").(string))

	if err != nil {
		if errAuth := new(auth.ErrorAuth); errors.As(err, &errAuth) {
			return ctx.Render("pages/error", map[string]any{
				"code":    errAuth.Code,
				"message": errAuth.Reason,
			})
		}

		log.Panic(err)
	}

	return ctx.Render("pages/profile", map[string]any{
		"user":      user,
		"lastLogin": utils.FormatDate(lastLoginMs),
		"total":     apiRes.TotalFiles,
	})
}

func HandleDetailDataClient(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)
	if !ok {
		log.Panic("invalid_user_struct")
	}

	agent := fiber.Get(fmt.Sprintf("%s/files/%s/%s", os.Getenv("API_SERVER_URL"), user.UserName, ctx.Params("name")))

	agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+ctx.Locals("token").(string))

	apiRes := new(models.DataFile)
	statusCode, body, errs := agent.Struct(&apiRes)
	if len(errs) > 0 {
		utils.Check(errs[0])
	}

	if statusCode != fiber.StatusOK {
		apiErr := new(models.ApiError)
		utils.Check(json.Unmarshal(body, &apiErr))

		return ctx.Render("pages/error", map[string]any{
			"code":    statusCode,
			"message": apiErr.Description,
		})
	}

	return ctx.Render("pages/details", map[string]any{
		"user": user,
		"file": struct {
			*models.DataFile
			UploadedAt   string
			UpdatedAt    string
			AutoDeleteAt string
		}{
			DataFile:     apiRes,
			UploadedAt:   utils.FormatDate(apiRes.UploadedAt),
			UpdatedAt:    utils.FormatDate(apiRes.UpdatedAt),
			AutoDeleteAt: utils.FormatDate(apiRes.AutoDeleteAt),
		},
		"type": "Update",
	})
}
