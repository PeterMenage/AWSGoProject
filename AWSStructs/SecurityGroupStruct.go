package AWSStructs

type SecurityGroup struct {
	SecurityGroups []struct {
		Description   string `json:"Description"`
		GroupID       string `json:"GroupId"`
		GroupName     string `json:"GroupName"`
		IPPermissions []struct {
			FromPort   int    `json:"FromPort"`
			IPProtocol string `json:"IpProtocol"`
			IPRanges   []struct {
				CidrIP string `json:"CidrIp"`
			} `json:"IpRanges"`
			ToPort int `json:"ToPort"`
		} `json:"IpPermissions"`
		IPPermissionsEgress []struct {
			IPProtocol string `json:"IpProtocol"`
			IPRanges   []struct {
				CidrIP string `json:"CidrIp"`
			} `json:"IpRanges"`
		} `json:"IpPermissionsEgress"`
		OwnerID string `json:"OwnerId"`
		VpcID   string `json:"VpcId"`
	} `json:"SecurityGroups"`
}
