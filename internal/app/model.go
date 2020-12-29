package app


type  Request struct {
	PatentID string `json:"patentId"`
	DemographicFactor Factor `json:"demographicFactor"`
	DiagnosedWithDiabetes bool `json:"diagnosedWithDiabetes"`
}
type Factor struct {
	Year int `json:"year"`
	Sex string `json:"sex"`
}


