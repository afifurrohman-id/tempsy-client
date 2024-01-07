package router

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/afifurrohman-id/tempsy-client/internal/client/models"
	"github.com/afifurrohman-id/tempsy-client/internal/client/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleUploadDashboardClient(ctx *fiber.Ctx) error {
	agent := fiber.Post(fmt.Sprintf("%s/files/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username")))

	agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+ctx.Locals("token").(string))
	agent.Body(ctx.Body())
	agent.Set(fiber.HeaderContentType, ctx.Get(fiber.HeaderContentType))
	agent.Set(utils.HeaderIsPublic, ctx.Get(utils.HeaderIsPublic))
	agent.Set(utils.HeaderPrivateUrlExpires, ctx.Get(utils.HeaderPrivateUrlExpires))
	agent.Set(utils.HeaderAutoDeletedAt, ctx.Get(utils.HeaderAutoDeletedAt))
	agent.Set(utils.HeaderFileName, ctx.Get(utils.HeaderFileName))

	apiRes := new(models.DataFile)
	statusCode, body, errs := agent.Struct(&apiRes)
	if len(errs) > 0 {
		log.Panic(errs[0])
	}

	if statusCode != fiber.StatusCreated {
		apiErr := new(models.ApiError)
		utils.Check(json.Unmarshal(body, &apiErr))

		return ctx.Status(statusCode).JSON(&apiErr)
	}

	return ctx.JSON(&apiRes)
}

func HandleUpdateDataClient(ctx *fiber.Ctx) error {
	agent := fiber.Put(fmt.Sprintf("%s/files/%s/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username"), ctx.Params("name")))

	agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+ctx.Locals("token").(string))
	agent.Body(ctx.Body())
	agent.Set(fiber.HeaderContentType, ctx.Get(fiber.HeaderContentType))
	agent.Set(utils.HeaderIsPublic, ctx.Get(utils.HeaderIsPublic))
	agent.Set(utils.HeaderPrivateUrlExpires, ctx.Get(utils.HeaderPrivateUrlExpires))
	agent.Set(utils.HeaderAutoDeletedAt, ctx.Get(utils.HeaderAutoDeletedAt))

	apiRes := new(models.DataFile)
	statusCode, body, errs := agent.Struct(&apiRes)
	if len(errs) > 0 {
		log.Panic(errs[0])
	}

	if statusCode != fiber.StatusOK {
		apiErr := new(models.ApiError)
		utils.Check(json.Unmarshal(body, &apiErr))

		return ctx.Status(statusCode).JSON(&apiErr)
	}

	return ctx.JSON(&apiRes)
}

func HandleDeleteDataClient(ctx *fiber.Ctx) error {
	agent := fiber.Delete(fmt.Sprintf("%s/files/%s/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username"), ctx.Params("name")))

	agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+ctx.Locals("token").(string))

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		utils.Check(errs[0])
	}

	if statusCode != fiber.StatusNoContent {
		apiRes := new(models.ApiError)
		utils.Check(json.Unmarshal(body, &apiRes))

		return ctx.Status(statusCode).JSON(&apiRes)
	}

	return ctx.SendStatus(statusCode)
}

func HandleDeleteAllDataClient(ctx *fiber.Ctx) error {
	agent := fiber.Delete(fmt.Sprintf("%s/files/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username")))

	agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+ctx.Locals("token").(string))

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		utils.Check(errs[0])
	}

	if statusCode != fiber.StatusNoContent {
		apiRes := new(models.ApiError)
		utils.Check(json.Unmarshal(body, &apiRes))

		return ctx.Status(statusCode).JSON(&apiRes)
	}

	return ctx.SendStatus(statusCode)
}
