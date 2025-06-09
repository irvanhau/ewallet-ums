package interfaces

import (
	"context"
	"ewallet-ums/external"
)

type IWallet interface {
	CreateWallet(ctx context.Context, userID uint) (*external.Wallet, error)
}
