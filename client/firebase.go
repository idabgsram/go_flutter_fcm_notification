package client

import (
	"errors"
	"os"

	"github.com/appleboy/go-fcm"
)

type FCM struct {
	Client *fcm.Client
}

func (fc *FCM) NewClient() (*fcm.Client, error) {
	apiKey := os.Getenv("FIREBASE_API_KEY")

	if apiKey == "" {
		return nil, errors.New("firebase api key required")
	}

	client, err := fcm.NewClient(apiKey)
	if err != nil {
		return nil, err
	}

	return client, nil
}
