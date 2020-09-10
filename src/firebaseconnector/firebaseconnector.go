// Copyright 2020 CBDT Technologies. All rights reserved.
// Use of this source code is governed by CBDT Technologies.
// Author Aravindh Chinnasamy Subburayar
package firebaseconnector

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

var client *db.Client
var ctx context.Context

func ConnectFireBase(url string) {

	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: url,
	}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}
	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

func IngestDataToFireBase(ref string, key string, value Visitor) {

	refs := client.NewRef(ref)
	usersRef := refs.Child(key)

	err := usersRef.Set(ctx, value)

	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
}

func UpdateDataToFireBase(ref string, key string, value string) {

}
