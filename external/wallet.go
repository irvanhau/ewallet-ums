package external

import (
	"bytes"
	"context"
	"encoding/json"
	"ewallet-ums/helpers"
	"net/http"

	"github.com/pkg/errors"
)

type Wallet struct {
	ID      int     `json:"id"`
	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

type ExtWallet struct {
}

func (*ExtWallet) CreateWallet(ctx context.Context, userID uint) (*Wallet, error) {
	req := Wallet{UserID: userID}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal json")
	}

	url := helpers.GetEnv("WALLET_HOST", "") + helpers.GetEnv("WALLET_ENDPOINT_CREATE", "")
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create wallet http request")
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.Errorf("got error response from wallet service: %d", resp.StatusCode)
	}

	result := &Wallet{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	defer resp.Body.Close()

	return result, nil

}
