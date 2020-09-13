// Copyright 2020 CBDT Technologies. All rights reserved.
// Use of this source code is governed by CBDT Technologies.
// Author Aravindh Chinnasamy Subburayar
package firestoreconnector

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"golang.org/x/exp/errors/fmt"
)

// export GOOGLE_APPLICATION_CREDENTIALS="/Users/dharshekthvel/Downloads/webcat-0346ea977d9a.json"

func init() {
	FireStoreConnectionPool()
}

func FireStoreConnectionPool() *firestore.Client {

	projectID := "webcat-284707"

	ctx := context.Background()

	FireStoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err, FireStoreClient)
	}

	return FireStoreClient
}

func IngestVisitorDataFireStore(FireStoreClient *firestore.Client, collectionName string, visitor VirtualVisitor) {

	projectID := "webcat-284707"

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	a, err := client.Collection(collectionName).Doc(visitor.FullName).Set(ctx, visitor)

	fmt.Println(a)
	if err != nil {
		log.Fatalf("Failed ingesting data : %v", err)
	}
}

//
func IngestExhibitorDataFireStore(FireStoreClient *firestore.Client, collectionName string, exhibitor VirtualExhibitor) {

	projectID := "webcat-284707"

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	a, err := client.Collection(collectionName).Doc(exhibitor.FullName).Set(ctx, exhibitor)

	fmt.Println(a)
	if err != nil {
		log.Fatalf("Failed ingesting data : %v", err)
	}
}

func UpdateDataFireStore() {

}
