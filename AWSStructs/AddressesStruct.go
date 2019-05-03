package AWSStructs

type Addresses struct {
	Addresses []struct {
		AllocationID   string `json:"AllocationId"`
		Domain         string `json:"Domain"`
		PublicIP       string `json:"PublicIp"`
		PublicIpv4Pool string `json:"PublicIpv4Pool"`
	} `json:"Addresses"`
}
