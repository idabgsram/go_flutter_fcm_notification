package main

import (
	"fmt"
	"guzram/utility/fcm_test/client"
	"guzram/utility/fcm_test/domain"
	"guzram/utility/fcm_test/utils"
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	fmt.Println("Initializing env ...")
	InitEnv()

	fmt.Println("Initializing fcm client ...")
	var fcm client.FCM
	fcmClient, errClient := fcm.NewClient()
	if errClient != nil {
		log.Fatal(errClient.Error())
	}

	fmt.Println("Sending message ...")
	msgBody := domain.FCMMessage{
		Title:   "Hello",
		Message: "This is a test message",
	}

	errSendMsg := utils.FirebaseSendPushNotification(fcmClient, msgBody)
	if errSendMsg != nil {
		log.Fatal(errSendMsg.Error())
	}

}
