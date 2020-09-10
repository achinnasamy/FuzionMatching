// Copyright 2020 CBDT Technologies. All rights reserved.
// Use of this source code is governed by CBDT Technologies.
// Author Aravindh Chinnasamy Subburayar
package firestoreconnector

import "golang.org/x/exp/errors/fmt"

func FuzionMatch() {

	productsMap := RetrieveExhibitorDataFireStore("exhibitor")

	vistorMap := RetrieveVisitorDataFireStore("vr_visitor")
	for userName, userProducts := range vistorMap {

		//fmt.Println("Name:", userName, "=>", "Element:", userProduct)

		for _, eachUserProduct := range userProducts {

			for companyName, companyProducts := range productsMap {

				for _, eachCompanyProduct := range companyProducts {

					//fmt.Println("Key:", companyName, "=>", "Element:", companyProduct)
					if eachUserProduct == eachCompanyProduct {
						fmt.Println("UserProduct:", eachUserProduct, "=>", "CompanyProduct:", eachCompanyProduct)
						fmt.Println("UserName:", userName, "=>", "Company Name:", companyName)
					}
				}

			}
		}
	}
}

func UpdateMatchingInFireStore() {

}
