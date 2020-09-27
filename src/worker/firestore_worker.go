package worker

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FuzionDataCarrier struct {
	UserName     string
	CompanyScore string
}

type FireStoreWorker struct {
}

func (firestoreWorker FireStoreWorker) InstantiateResources() {
	return
}

func (firestoreWorker FireStoreWorker) WorkersToCreate() int {
	return 2
}

func (firestoreWorker FireStoreWorker) HeavyWork(done chan int) bool {

	//inputURL := "https://webcat-284707.firebaseio.com/"

	//firebaseconnector.ConnectFireBase(inputURL)

	for i := 0; i < 999999; i++ {
		time.Sleep(100 * time.Millisecond)
		//fmt.Println("Matching running on firestore")

		hitFireStore()
		// v := firebaseconnector.Visitor{
		// 	DateOfBirth: "June 24, 1912",
		// 	FullName:    "Stallman",
		// 	CompanyName: "Alphanet",
		// 	Product:     "Linen",
		// }
		//firebaseconnector.IngestDataToFireBase("visitor_details", "20001996", v)
	}

	if done != nil {
		done <- 0 // Signal that we're done
	}

	return true
}

func hitFireStore() {
	resultantArray := FuzionMatch()

	for _, each := range resultantArray {
		if each.CompanyScore != "" {
			fmt.Println("UserName:", each.UserName, "=>", "value - ", each.CompanyScore)
			UpdateFireStore("vr_visitor", each.UserName, each.CompanyScore)
		}
	}

	fmt.Println("END")

}

func FuzionMatch() []FuzionDataCarrier {

	var howManyRecommendations int = 0
	var resultArray []FuzionDataCarrier
	var _companyScore string = ""

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
						fmt.Println("UserProduct:", eachUserProduct, "=>", "CompanyProduct:", eachCompanyProduct)
						//fmt.Println("UserName:", userName, "=>", "Company Name:", companyName)
						matchingCounter++
					}

					totalProducts++
					//fmt.Println(len(companyProducts))
				}

				percentage := (matchingCounter * 100 / totalProducts)

				if percentage > 0 {
					if howManyRecommendations >= 0 {
						_companyScore = _companyScore + " " + companyName + ":" + strconv.Itoa(percentage)
					}
					howManyRecommendations++
					//_companyScore = _companyScore + " " + companyName + ":" + strconv.Itoa(percentage)

					//resultArray = append(resultArray, FuzionDataCarrier{userName, companyName + " : " + strconv.Itoa(percentage)})
				}

				matchingCounter = 0
				totalProducts = 0

			}
		}

		resultArray = append(resultArray, FuzionDataCarrier{userName, SortCompaniesInReverseBasedOnScore(_companyScore)})
		_companyScore = ""
		howManyRecommendations = 0
	}

	return resultArray

}

func SortCompaniesInReverseBasedOnScore(companyScore string) string {

	var _finalCompanyScore string = ""
	resultantMap := make(map[int]string)

	splitOnCompany := strings.Split(companyScore, " ")
	for _, eachCompany := range splitOnCompany {

		if eachCompany != "" {
			companyAndScore := strings.Split(strings.TrimSpace(eachCompany), ":")
			integer_value, _ := strconv.Atoi(companyAndScore[1])
			resultantMap[integer_value] = companyAndScore[0]
		}
	}

	keys := make([]int, len(resultantMap))
	i := 0
	for k := range resultantMap {
		keys[i] = k
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, v := range keys {
		_finalCompanyScore = _finalCompanyScore + resultantMap[v] + ":" + strconv.Itoa(v) + " "
	}

	return _finalCompanyScore
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

func UpdateFireStore(collectionName string, documentToBeUpdated string, data string) {

	projectID := "webcat-284707"

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	result, err := client.Collection(collectionName).Doc(documentToBeUpdated).Set(ctx, map[string]interface{}{
		"Score": data,
	}, firestore.MergeAll)

	if err != nil {
		log.Printf("An error has occurred: %s", err, result)
	}

}
