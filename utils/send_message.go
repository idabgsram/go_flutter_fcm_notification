package utils

import (
	"errors"
	"fmt"
	"guzram/utility/fcm_test/domain"
	"os"

	"github.com/appleboy/go-fcm"
)

func FirebaseSendPushNotification(fc *fcm.Client, body domain.FCMMessage) error {

	// deviceToken originally is a slice, but in this case,
	// we will use a single token from env and add it to a slice
	deviceToken := os.Getenv("FIREBASE_DEVICE_TOKEN_ID")
	if deviceToken == "" {
		return errors.New("firebase device token required")
	}

	devicesToken := []string{deviceToken}

	msg := &fcm.Message{
		RegistrationIDs: devicesToken,
		Data: map[string]interface{}{
			"click_action": "FLUTTER_NOTIFICATION_CLICK",
			"page":         "home",
			"id":           1,
		},
		Priority: "high",
		Notification: &fcm.Notification{
			Title: body.Title,
			Body:  body.Message,
		},
	}

	response, err := fc.Send(msg)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", response)

	return nil
}
