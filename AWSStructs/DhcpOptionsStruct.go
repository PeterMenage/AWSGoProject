type DhcpOptions struct {
	DhcpOptions []struct {
		DhcpConfigurations []struct {
			Key    string `json:"Key"`
			Values []struct {
				Value string `json:"Value"`
			} `json:"Values"`
		} `json:"DhcpConfigurations"`
		DhcpOptionsID string `json:"DhcpOptionsId"`
		OwnerID       string `json:"OwnerId"`
	} `json:"DhcpOptions"`
}