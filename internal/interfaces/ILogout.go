package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ILogoutHandler interface {
	Logout(c *gin.Context)
}

type ILogoutService interface {
	Logout(ctx context.Context, token string) error
}
