package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// MakeAuthLoginHandler makes a handler which redirects the user to the GitHub
// OAuth page
func MakeAuthLoginHandler(ghClientID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		redirectURL, err := url.Parse("https://github.com/login/oauth/authorize")

		if err != nil {
			panic(fmt.Sprintf("error building GitHub OAuth redirect URL: %s",
				err.Error()))
		}

		queryParams := redirectURL.Query()
		queryParams.Set("client_id", ghClientID)

		redirectURL.RawQuery = queryParams.Encode()

		c.Redirect(http.StatusMovedPermanently, redirectURL.String())
	}
}

// MakeAuthCallbackHandler makes a handler which exchanges a Github temporary auth code for a longer living Github
// auth token
func MakeAuthCallbackHandler(ghClientID string, ghClientSecret string) gin.HandlerFunc {
	type callbackReq struct {
		Code string
	}

	return func(c *gin.Context) {
		var body callbackReq

		err := c.ShouldBind(&body)
		if err != nil {
			panic(fmt.Sprintf("request not formatted correctly: %s", err.Error()))
		}

		c.JSON(http.StatusOK, gin.H{
			"code": body.Code,
		})
	}
}
