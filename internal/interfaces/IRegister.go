package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

type IRegisterRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}

type IRegisterService interface {
	Register(ctx context.Context, req models.User) (interface{}, error)
}
