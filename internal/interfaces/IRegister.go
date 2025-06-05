package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterHandler interface {
	Register(c *gin.Context)
}

type IRegisterService interface {
	Register(ctx context.Context, req models.User) (interface{}, error)
}
