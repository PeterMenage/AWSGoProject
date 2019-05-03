package AWSStructs

type NatGateways struct {
	NatGateways []struct {
		CreateTime          string `json:"CreateTime"`
		NatGatewayAddresses []struct {
			AllocationID       string `json:"AllocationId"`
			NetworkInterfaceID string `json:"NetworkInterfaceId"`
			PrivateIP          string `json:"PrivateIp"`
			PublicIP           string `json:"PublicIp"`
		} `json:"NatGatewayAddresses"`
		NatGatewayID string `json:"NatGatewayId"`
		State        string `json:"State"`
		SubnetID     string `json:"SubnetId"`
		VpcID        string `json:"VpcId"`
	} `json:"NatGateways"`
}