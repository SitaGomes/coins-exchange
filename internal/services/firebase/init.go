package firebase

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Firebase struct {
	app    *firebase.App
	client *firestore.Client
}

func (f *Firebase) initialize() error {
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("FIREBASE_ADMIN_CONFIG"))

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}

	f.app = app
	f.client = client
	return nil
}

func GetInstance() *Firebase {
	instance := &Firebase{}
	if err := instance.initialize(); err != nil {
		log.Fatal("Failed to initialize Firebase ", err)
		return nil
	}

	return instance
}

func (f *Firebase) GetFirestoreClient() *firestore.Client {
	return GetInstance().client
}
