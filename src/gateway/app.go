// Copyright 2020 CBDT Technologies. All rights reserved.
// Use of this source code is governed by CBDT Technologies.
//
// Main entrypoint of the application
// Author Aravindh Chinnasamy Subburayar
package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"worker"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

//
//Main entrypoint of the walkover application
//
func main() {
	// start application
	fmt.Println("Application Started !!!")

	hitFireStore()

	// End application
	fmt.Println("Application Ended. !!!")
}

func hitFirebase() {

	// inputURL := "https://webcat-284707.firebaseio.com/"

	// firebaseconnector.ConnectFireBase(inputURL)
	// firebaseconnector.IngestDataToFireBase("", "", "")

	worker.StartWorking(worker.FirebaseWorker{})
}

func hitFireStore() {
	//firestoreconnector.FuzionMatch()

	productsMap := RetrieveExhibitorDataFireStore("exhibitor")
	vistorMap := RetrieveVisitorDataFireStore("vr_visitor")

	var totalProducts int = 0
	var matchingCounter int = 0

	for userName, userProducts := range vistorMap {

		//fmt.Println("Name:", userName, "=>", "Element:", userProduct)

		for _, eachUserProduct := range userProducts {

			for companyName, companyProducts := range productsMap {

				for _, eachCompanyProduct := range companyProducts {

					//fmt.Println("Key:", companyName, "=>", "Element:", eachCompanyProduct)
					if eachUserProduct == eachCompanyProduct {
						//fmt.Println("UserProduct:", eachUserProduct, "=>", "CompanyProduct:", eachCompanyProduct)
						//fmt.Println("UserName:", userName, "=>", "Company Name:", companyName)
						matchingCounter++
					}

					totalProducts++
					//fmt.Println(len(companyProducts))
				}

				percentage := (matchingCounter * 100 / totalProducts)

				if percentage > 0 {
					//fmt.Println(percentage)
					fmt.Println("UserName:", userName, "=>", "Company Name:", companyName, " => ", "percentage", percentage)
				}

				matchingCounter = 0
				totalProducts = 0

			}
		}
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
