package connect

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func Connect() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("fir-tutorial-a0fcd-firebase-adminsdk-bjopb-d6e57eae98.json")

	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		panic("Gagal Inisialisasi Database Firebase : " + err.Error())
	}

	client, errClient := app.Firestore(ctx)

	if errClient != nil {
		panic("Gagal Terkoneksi Ke Database Firebase : " + errClient.Error())
	}

	FirestoreClient = client

	log.Println("Berhasil Terhubung Ke Database 🐦‍🔥")

}
