package fcm

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

/**
	JSON file that contains your service account key
	export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/[FILE_NAME].json"
**/

type App struct {
	firebaseApp     *firebase.App
	messagingClient *messaging.Client
}

func NewApp() (App, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return App{}, err
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
		return App{}, err
	}

	return App{
		firebaseApp:     app,
		messagingClient: client,
	}, err
}

func (app App) SendMessage(
	registrationToken string,
	title string, body string, data map[string]string) error {

	if app.firebaseApp == nil {
		return fmt.Errorf("Cannot send message, failed at initialization")
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: registrationToken,
	}

	if data != nil {
		message.Data = data
	}

	response, err := app.messagingClient.Send(context.Background(), message)
	if err != nil {
		log.Printf("error sending message: %v\n", err)
		return err
	}

	log.Print("Successfully sent message:", response)
	return nil
}
