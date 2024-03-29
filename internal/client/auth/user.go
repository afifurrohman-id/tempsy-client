package auth

import (
	"os"
	"strings"

	"github.com/afifurrohman-id/tempsy-client/internal/client/models"
	"github.com/afifurrohman-id/tempsy-client/internal/client/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(token string) (*models.User, error) {
	agent := fiber.Get(os.Getenv("API_SERVER_URL") + "/auth/userinfo/me")

	agent.Timeout(utils.DefaultApiTimeout)
	agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+token)

	apiRes := new(models.User)
	statusCode, body, errs := agent.Struct(&apiRes)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return nil, &ErrorAuth{
			Code:   statusCode,
			Reason: string(body),
		}
	}

	if apiRes.Picture == "" && strings.HasPrefix(apiRes.UserName, GuestUsernamePrefix) {
		apiRes.Picture = "https://placehold.co/96x96/png?text=a"
	}

	return apiRes, nil
}
