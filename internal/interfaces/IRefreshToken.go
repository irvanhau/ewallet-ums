package interfaces

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRefreshTokenHandler interface {
	RefreshToken(c *gin.Context)
}

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}
