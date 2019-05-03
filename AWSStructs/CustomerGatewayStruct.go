package AWSStructs

type CustomerGateway struct {
	CustomerGateways []struct {
		BgpAsn            string `json:"BgpAsn"`
		CustomerGatewayID string `json:"CustomerGatewayId"`
		IPAddress         string `json:"IpAddress"`
		State             string `json:"State"`
		Tags              []struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
		} `json:"Tags"`
		Type string `json:"Type"`
	} `json:"CustomerGateways"`
}
