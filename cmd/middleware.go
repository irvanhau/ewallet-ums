package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(c *gin.Context) {
	auth := c.Request.Header.Get("authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(c.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	claim, err := helpers.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	c.Set("token", claim)

	c.Next()
}
