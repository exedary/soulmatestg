package auth

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/exedary/soulmates/internal/domain/person"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

const (
	resourceName       = "/auth"
	sessionStateToken  = "session_state_token"
	defaultSessionName = "default"
	persistedSession   = "persisted"
)

var Module = fx.Invoke(Register)

type controller struct {
	personRepository person.Repository
}

func Register(router *gin.RouterGroup, repository person.Repository) {
	controller := &controller{personRepository: repository}
	router.GET(resourceName+"/google/login", controller.login)
	router.GET(resourceName+"/google/callback", controller.processGoogleCallback)
}

func (controller *controller) login(c *gin.Context) {
	stateHash := randomString(30)

	url := SignInWithGoogle(c.Request.Context(), controller.personRepository, stateHash)

	//TO-DO: Correctly add cookies to response in order to secure them
	c.SetCookie(sessionStateToken, stateHash, 3600, "", "localhost", false, true)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (controller *controller) processGoogleCallback(c *gin.Context) {
	authCode := c.Query("code")
	oauthState, _ := c.Cookie(sessionStateToken)

	if oauthState != c.Query("state") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ProcessGoogleCallback(c.Request.Context(), controller.personRepository, authCode)

	c.JSON(http.StatusOK, nil)
}

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
