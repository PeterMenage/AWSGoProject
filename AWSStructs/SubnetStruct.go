package AWSStructs

type Subnet struct {
	Subnets []struct {
		AssignIpv6AddressOnCreation bool   `json:"AssignIpv6AddressOnCreation"`
		AvailabilityZone            string `json:"AvailabilityZone"`
		AvailabilityZoneID          string `json:"AvailabilityZoneId"`
		AvailableIPAddressCount     int    `json:"AvailableIpAddressCount"`
		CidrBlock                   string `json:"CidrBlock"`
		DefaultForAz                bool   `json:"DefaultForAz"`
		MapPublicIPOnLaunch         bool   `json:"MapPublicIpOnLaunch"`
		OwnerID                     string `json:"OwnerId"`
		State                       string `json:"State"`
		SubnetArn                   string `json:"SubnetArn"`
		SubnetID                    string `json:"SubnetId"`
		VpcID                       string `json:"VpcId"`
	} `json:"Subnets"`
}
