package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kzdv/api/internal/v1/dto"
	"github.com/kzdv/api/pkg/config"
	"github.com/kzdv/api/pkg/gin/response"
	"github.com/kzdv/api/pkg/oauth"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Login to account
// @Summary Login to account
// @Tags user, oauth
// @Param redirect path string false "Redirect URL"
// @Success 307
// @Failure 500 {object} response.R
// @Router /v1/user/login [GET]
func GetLogin(c *gin.Context) {
	state, _ := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 64)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Set("redirect", c.Query("redirect"))
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, oauth.OAuthConfig.AuthCodeURL(state))
}

// Login callback
// @Summary Login callback
// @Tags user, oauth
// @Success 307
// @Success 200 {object} response.R
// @Failure 400 {object} response.R
// @Failure 403 {object} response.R
// @Failure 500 {object} response.R
// @Router /v1/user/login/callback [GET]
func GetLoginCallback(c *gin.Context) {
	session := sessions.Default(c)
	state := session.Get("state")
	if state == nil {
		response.RespondError(c, http.StatusForbidden, "Forbidden")
		return
	}
	if state != c.Query("state") {
		response.RespondError(c, http.StatusForbidden, "Forbidden")
		return
	}
	token, err := oauth.OAuthConfig.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, "Bad Request")
		return
	}
	res, err := http.NewRequest("GET", fmt.Sprintf("%s%s", config.Cfg.OAuth.BaseURL, config.Cfg.OAuth.Endpoints.UserInfo), nil)
	res.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	res.Header.Add("Accept", "application/json")
	res.Header.Add("User-Agent", "kzdv-api")
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer res.Body.Close()
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	user := &dto.SSOUserResponse{}
	if err := json.Unmarshal(contents, &user); err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	session.Delete("state")
	session.Set("cid", fmt.Sprint(user.User.CID))
	session.Save()

	redirect := session.Get("redirect")
	if redirect != nil {
		c.Redirect(http.StatusTemporaryRedirect, redirect.(string))
		return
	}
	response.RespondMessage(c, http.StatusOK, "Logged In")
}
