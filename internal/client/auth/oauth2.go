package auth

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/afifurrohman-id/tempsy-client/internal/client/models"
	"github.com/afifurrohman-id/tempsy-client/internal/client/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var ErrorGOAuth2 = errors.New("oauth2_error_response_code_not_ok")

type GOAuth2Config struct {
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	CallbackURL  string   `json:"callbackUrl"`
	Scopes       []string `json:"scopes"`
}

func New() (*GOAuth2Config, error) {
	oConfJSON := os.Getenv("GOOGLE_OAUTH2_CONFIG")

	oConf := new(GOAuth2Config)

	if err := json.Unmarshal([]byte(oConfJSON), &oConf); err != nil {
		return nil, err
	}

	return oConf, nil
}

func GetGoogleAccountInfo(accessToken string) (*models.GoogleAccountInfo, error) {
	const timeout = 10 * time.Second

	agent := fiber.Get("https://www.googleapis.com/userinfo/v2/me")

	agent.Set(fiber.HeaderAuthorization, utils.BearerPrefix+accessToken)
	// TODO: Add parameter timeout
	agent.Timeout(timeout)

	userinfo := new(models.GoogleAccountInfo)

	statusCode, body, errs := agent.Struct(&userinfo)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if statusCode != fiber.StatusOK {
		log.Errorf("response_from_%d_body_%s", statusCode, string(body))
		return nil, ErrorGOAuth2
	}

	userinfo.UserName = strings.ReplaceAll(strings.Join(strings.SplitN(userinfo.Email, "@", 2), "-"), ".", "-")

	return userinfo, nil
}
