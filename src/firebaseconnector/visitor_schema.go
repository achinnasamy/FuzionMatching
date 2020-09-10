// Copyright 2020 CBDT Technologies. All rights reserved.
// Use of this source code is governed by CBDT Technologies.
// Author Aravindh Chinnasamy Subburayar
package firebaseconnector

type Visitor struct {
	DateOfBirth   string `json:"date_of_birth,omitempty"`
	FullName      string `json:"full_name,omitempty"`
	MatchingScore string `json:"matching_score,omitempty"`
	CompanyName   string `json:"company_name,omitempty"`
	Product       string `json:"product,omitempty"`
}
