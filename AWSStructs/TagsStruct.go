package AWSStructs


type Tags struct {
	Tags []struct {
		Key          string `json:"Key"`
		ResourceID   string `json:"ResourceId"`
		ResourceType string `json:"ResourceType"`
		Value        string `json:"Value"`
	} `json:"Tags"`
}