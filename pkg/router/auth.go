package router

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/afifurrohman-id/tempsy-client/internal/client/auth"
	"github.com/afifurrohman-id/tempsy-client/internal/client/models"
	"github.com/afifurrohman-id/tempsy-client/internal/client/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Six months in seconds
const maxAgeCookie = 6 * 30 * 24 * 60 * 60

func OAuth2Callback(ctx *fiber.Ctx) error {
	oAuth2, err := auth.New()
	utils.Check(err)

	tokens, err := oAuth2.ExchangeCode(ctx.Query("code"))
	if err != nil {
		if errAuth := new(auth.ErrorAuth); errors.As(err, &errAuth) {
			return ctx.Redirect("/auth/login")
		}
		log.Panic(err)
	}

	user, err := auth.GetGoogleAccountInfo(tokens.AccessToken)
	utils.Check(err)

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		MaxAge:   maxAgeCookie,
		Secure:   os.Getenv("APP_ENV") == "production",
		HTTPOnly: os.Getenv("APP_ENV") == "production",
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "last_login",
		Value:    fmt.Sprintf("%d", time.Now().UnixMilli()),
		Path:     "/",
		MaxAge:   maxAgeCookie,
		Secure:   os.Getenv("APP_ENV") == "production",
		HTTPOnly: os.Getenv("APP_ENV") == "production",
	})

	return ctx.Redirect("/dashboard/" + user.UserName)
}

func AuthLogin(ctx *fiber.Ctx) error {
	token := ctx.Cookies("token")
	oAuth2, err := auth.New()
	utils.Check(err)

	if _, err := oAuth2.AccessToken(token); err == nil {
		return ctx.Redirect("/")
	}

	if _, err = auth.GetUserInfo(token); err == nil {
		return ctx.Redirect("/")
	}

	if ctx.Query("type", "oauth2") == "guest" {

		agent := fiber.Get(os.Getenv("API_SERVER_URL") + "/auth/guest/token")

		agent.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSON)

		apiRes := new(models.GuestToken)
		statusCode, body, errs := agent.Struct(&apiRes)
		if len(errs) > 0 {
			log.Panic(errs[0])
		}

		if statusCode != fiber.StatusOK {
			if statusCode == fiber.StatusBadRequest {
				return ctx.Redirect("/")
			}
			return ctx.Render("pages/error", map[string]any{
				"code":    statusCode,
				"message": string(body),
			})
		}

		ctx.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    apiRes.AccessToken,
			Path:     "/",
			MaxAge:   apiRes.ExpiresIn, // convert to seconds
			Secure:   os.Getenv("APP_ENV") == "production",
			HTTPOnly: os.Getenv("APP_ENV") == "production",
		})

		ctx.Cookie(&fiber.Cookie{
			Name:     "last_login",
			Value:    fmt.Sprintf("%d", time.Now().UnixMilli()),
			Path:     "/",
			MaxAge:   apiRes.ExpiresIn, // convert to seconds
			Secure:   os.Getenv("APP_ENV") == "production",
			HTTPOnly: os.Getenv("APP_ENV") == "production",
		})

		apiResUser, err := auth.GetUserInfo(apiRes.AccessToken)
		if err != nil {
			if errAuth := new(auth.ErrorAuth); errors.As(err, &errAuth) {
				return ctx.Render("pages/error", map[string]any{
					"code":    errAuth.Code,
					"message": errAuth.Reason,
				})
			}
			log.Panic(err)
		}

		return ctx.Redirect("/dashboard/" + apiResUser.UserName)
	}

	return ctx.Redirect(oAuth2.RedirectUrl())
}

func AuthLogout(ctx *fiber.Ctx) error {
	oAuth2, err := auth.New()
	utils.Check(err)

	tokens, err := oAuth2.AccessToken(ctx.Cookies("token")) // refresh token if oauth2 or access token if guest
	if err != nil {
		if errAuth := new(auth.ErrorAuth); !errors.As(err, &errAuth) {
			log.Panic(err)
		}
	} else {
		if err = oAuth2.RevokeAccessToken(tokens.AccessToken); err != nil {
			log.Error(err)
		}
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Expires:  time.Now().Add(-3 * time.Second), // 3 seconds ago, to delete cookie
		Secure:   os.Getenv("APP_ENV") == "production",
		HTTPOnly: os.Getenv("APP_ENV") == "production",
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "last_login",
		Path:     "/",
		Expires:  time.Now().Add(-3 * time.Second),
		Secure:   os.Getenv("APP_ENV") == "production",
		HTTPOnly: os.Getenv("APP_ENV") == "production",
	})

	return ctx.Redirect("/")
}
