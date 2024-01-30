package middleware

import (
	"errors"
	"strings"

	"github.com/afifurrohman-id/tempsy-client/internal/client/auth"
	"github.com/afifurrohman-id/tempsy-client/internal/client/models"
	"github.com/afifurrohman-id/tempsy-client/internal/client/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CheckAuth(ctx *fiber.Ctx) error {
	var (
		path  = ctx.Path()
		user  = new(models.User)
		token = ctx.Cookies("token") // refresh token if oauth2 or access token if guest
	)
	o2, err := auth.New()
	utils.Check(err)

	tokens, err := o2.AccessToken(token)
	if err != nil {
		// try to get guest user
		if errAuth := new(auth.ErrorAuth); errors.As(err, &errAuth) {
			apiRes, err := auth.GetUserInfo(token)
			if err != nil {
				if errors.As(err, &errAuth) {
					if path == "/" {
						return ctx.Next()
					}
					return ctx.Redirect("/auth/login")
				}
				log.Panic(err)
			}

			if !strings.Contains(apiRes.UserName, auth.GuestUsernamePrefix) {
				// for root path it will render no user.
				if path == "/" {
					return ctx.Next()
				}

				return ctx.Redirect("/auth/login")
			}

			user = apiRes
			user.Picture = "https://placehold.co/96x96/png?text=a"
		} else {
			log.Panic(err)
		}
	} else {
		userInfo, err := auth.GetGoogleAccountInfo(tokens.AccessToken)
		utils.Check(err)

		if !userInfo.VerifiedEmail {
			if path == "/" {
				return ctx.Next()
			}
			return ctx.Redirect("/auth/login")
		}

		user = userInfo.User

		// change to access token
		token = tokens.AccessToken
	}
	ctx.Locals("user", user)
	ctx.Locals("token", token) // access token
	if path != "/" && user.UserName != ctx.Params("username") {
		return ctx.Redirect("/dashboard/" + user.UserName)
	}
	return ctx.Next()
}

// SetRealIpClient TODO: It's Really work?
func SetRealIpClient(ctx *fiber.Ctx) error {
	ctx.Request().Header.Set(utils.HeaderRealIp, ctx.IP())
	return ctx.Next()
}
