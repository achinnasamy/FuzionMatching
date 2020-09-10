package firebaseconnector

type Exhibitor struct {
	DateOfBirth   string `json:"date_of_birth,omitempty"`
	FullName      string `json:"full_name,omitempty"`
	MatchingScore string `json:"matching_score,omitempty"`
	CompanyName   string `json:"company_name,omitempty"`
	Product       string `json:"product,omitempty"`
}
