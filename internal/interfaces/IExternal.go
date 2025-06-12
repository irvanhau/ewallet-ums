package interfaces

import (
	"context"
	"ewallet-ums/external"
)

type IExternal interface {
	CreateWallet(ctx context.Context, userID uint) (*external.Wallet, error)
	SendNotification(ctx context.Context, recipient string, templateName string, placeholder map[string]string) error
}
