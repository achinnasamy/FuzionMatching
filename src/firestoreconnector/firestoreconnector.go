package firestoreconnector

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"golang.org/x/exp/errors/fmt"
	"google.golang.org/api/iterator"
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

func RetrieveExhibitorDataFireStore(collectionName string) map[string][]string {

	var intermittentStorage map[string][]string = make(map[string][]string)
	var arrayStore []string

	projectID := "webcat-284707"

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		//intermittentStorage[doc.Data()["FullName"].(string)] = doc.Data()["Product"].(map[string][]string)

		//fmt.Println(doc.Data()["Product"])

		for topkey, value := range doc.Data()["Product"].(map[string]interface{}) {

			str := fmt.Sprintf("%v", value)

			arrayStore = append(arrayStore, topkey)

			for _, value := range strings.Split(str, " ") {

				arrayStore = append(arrayStore, value)
			}
		}

		intermittentStorage[doc.Data()["FullName"].(string)] = arrayStore
		arrayStore = nil
	}

	return intermittentStorage
}

func RetrieveVisitorDataFireStore(collectionName string) map[string][]string {

	var intermittentStorage map[string][]string = make(map[string][]string)
	var arrayStore []string

	projectID := "webcat-284707"

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		products := fmt.Sprintf("%v", doc.Data()["Product"])

		for _, value := range strings.Split(products, " ") {

			arrayStore = append(arrayStore, value)
		}

		intermittentStorage[doc.Data()["FullName"].(string)] = arrayStore
		arrayStore = nil
	}

	return intermittentStorage
}

func UpdateDataFireStore() {

}
