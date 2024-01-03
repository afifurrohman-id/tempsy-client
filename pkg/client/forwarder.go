package client

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/afifurrohman-id/tempsy-client/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleUploadDashboardClient(ctx *fiber.Ctx) error {
	agent := fiber.Post(fmt.Sprintf("%s/files/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username")))

	agent.Set(fiber.HeaderAuthorization, internal.BearerPrefix+ctx.Locals("token").(string))
	agent.Body(ctx.Body())

	agent.Set(fiber.HeaderContentType, ctx.Get(fiber.HeaderContentType))
	agent.Set(internal.HeaderIsPublic, ctx.Get(internal.HeaderIsPublic))
	agent.Set(internal.HeaderPrivateUrlExpires, ctx.Get(internal.HeaderPrivateUrlExpires))
	agent.Set(internal.HeaderAutoDeletedAt, ctx.Get(internal.HeaderAutoDeletedAt))
	agent.Set(internal.HeaderFileName, ctx.Get(internal.HeaderFileName))

	apiRes := new(internal.DataFile)
	statusCode, body, errs := agent.Struct(&apiRes)
	if len(errs) > 0 {
		log.Panic(errs[0])
	}

	if statusCode != fiber.StatusCreated {
		apiErr := new(internal.ApiError)
		internal.Check(json.Unmarshal(body, &apiErr))

		return ctx.Status(statusCode).JSON(&apiErr)
	}

	return ctx.JSON(&apiRes)
}

func HandleUpdateDataClient(ctx *fiber.Ctx) error {
	agent := fiber.Put(fmt.Sprintf("%s/files/%s/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username"), ctx.Params("name")))

	agent.Set(fiber.HeaderAuthorization, internal.BearerPrefix+ctx.Locals("token").(string))
	agent.Body(ctx.Body())
	agent.Set(fiber.HeaderContentType, ctx.Get(fiber.HeaderContentType))
	agent.Set(internal.HeaderIsPublic, ctx.Get(internal.HeaderIsPublic))
	agent.Set(internal.HeaderPrivateUrlExpires, ctx.Get(internal.HeaderPrivateUrlExpires))
	agent.Set(internal.HeaderAutoDeletedAt, ctx.Get(internal.HeaderAutoDeletedAt))

	apiRes := new(internal.DataFile)
	statusCode, body, errs := agent.Struct(&apiRes)
	if len(errs) > 0 {
		log.Panic(errs[0])
	}

	if statusCode != fiber.StatusOK {
		apiErr := new(internal.ApiError)
		internal.Check(json.Unmarshal(body, &apiErr))

		return ctx.Status(statusCode).JSON(&apiErr)
	}

	return ctx.JSON(&apiRes)
}

func HandleDeleteDataClient(ctx *fiber.Ctx) error {
	agent := fiber.Delete(fmt.Sprintf("%s/files/%s/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username"), ctx.Params("name")))

	agent.Set(fiber.HeaderAuthorization, internal.BearerPrefix+ctx.Locals("token").(string))

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		internal.Check(errs[0])
	}

	if statusCode != fiber.StatusNoContent {
		apiRes := new(internal.ApiError)
		internal.Check(json.Unmarshal(body, &apiRes))

		return ctx.Status(statusCode).JSON(&apiRes)
	}

	return ctx.SendStatus(statusCode)
}

func HandleDeleteAllDataClient(ctx *fiber.Ctx) error {
	agent := fiber.Delete(fmt.Sprintf("%s/files/%s", os.Getenv("API_SERVER_URL"), ctx.Params("username")))
	agent.Set(fiber.HeaderAuthorization, internal.BearerPrefix+ctx.Locals("token").(string))

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		internal.Check(errs[0])
	}

	if statusCode != fiber.StatusNoContent {
		apiRes := new(internal.ApiError)
		internal.Check(json.Unmarshal(body, &apiRes))

		return ctx.Status(statusCode).JSON(&apiRes)
	}

	return ctx.SendStatus(statusCode)
}
