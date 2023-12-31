package middleware

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/afifurrohman-id/tempsy-client/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CheckAuth(ctx *fiber.Ctx) error {
	var (
		path  = ctx.Path()
		user  = new(internal.User)
		token = ctx.Cookies("token") // refresh token if oauth2 or access token if guest
	)
	o2, err := internal.New()
	internal.Check(err)

	tokens, err := o2.AccessToken(token)

	// TODO:Validate if user is authorized
	if err != nil {
		// try to get guest user
		if errors.Is(err, internal.ErrorGOAuth2) {
			agent := fiber.Get(fmt.Sprintf("%s/auth/userinfo/me", os.Getenv("API_SERVER_URI")))
			agent.Set(fiber.HeaderAuthorization, internal.BearerPrefix+token)
			apiRes := new(internal.User)

			agent.Timeout(10 * time.Second)
			statusCode, _, errs := agent.Struct(&apiRes)
			if len(errs) > 0 {
				log.Panic(errs[0])
			}

			if statusCode != fiber.StatusOK || !strings.HasPrefix(apiRes.UserName, internal.GuestUsernamePrefix) {
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
		userInfo, err := internal.GetGoogleAccountInfo(tokens.AccessToken)
		internal.Check(err)

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
	ctx.Request().Header.Set(internal.HeaderRealIp, ctx.IP())
	return ctx.Next()
}
