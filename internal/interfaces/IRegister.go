package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

type IRegisterService interface {
	Register(ctx context.Context, req models.User) (interface{}, error)
}
