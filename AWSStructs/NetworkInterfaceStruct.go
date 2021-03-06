package AWSStructs

type AutoGenerated struct {
	NetworkInterfaces []struct {
		Attachment struct {
			AttachmentID        string `json:"AttachmentId"`
			DeleteOnTermination bool   `json:"DeleteOnTermination"`
			DeviceIndex         int    `json:"DeviceIndex"`
			InstanceOwnerID     string `json:"InstanceOwnerId"`
			Status              string `json:"Status"`
		} `json:"Attachment,omitempty"`
		AvailabilityZone string `json:"AvailabilityZone"`
		Description      string `json:"Description"`
		Groups           []struct {
			GroupID   string `json:"GroupId"`
			GroupName string `json:"GroupName"`
		} `json:"Groups"`
		InterfaceType      string `json:"InterfaceType"`
		MacAddress         string `json:"MacAddress"`
		NetworkInterfaceID string `json:"NetworkInterfaceId"`
		OwnerID            string `json:"OwnerId"`
		PrivateDNSName     string `json:"PrivateDnsName"`
		PrivateIPAddress   string `json:"PrivateIpAddress"`
		PrivateIPAddresses []struct {
			Primary          bool   `json:"Primary"`
			PrivateDNSName   string `json:"PrivateDnsName"`
			PrivateIPAddress string `json:"PrivateIpAddress"`
		} `json:"PrivateIpAddresses"`
		RequesterID      string `json:"RequesterId,omitempty"`
		RequesterManaged bool   `json:"RequesterManaged"`
		SourceDestCheck  bool   `json:"SourceDestCheck"`
		Status           string `json:"Status"`
		SubnetID         string `json:"SubnetId"`
		VpcID            string `json:"VpcId"`
		Association      struct {
			IPOwnerID     string `json:"IpOwnerId"`
			PublicDNSName string `json:"PublicDnsName"`
			PublicIP      string `json:"PublicIp"`
		} `json:"Association,omitempty"`
		Attachment struct {
			AttachTime          string `json:"AttachTime"`
			AttachmentID        string `json:"AttachmentId"`
			DeleteOnTermination bool   `json:"DeleteOnTermination"`
			DeviceIndex         int    `json:"DeviceIndex"`
			InstanceID          string `json:"InstanceId"`
			InstanceOwnerID     string `json:"InstanceOwnerId"`
			Status              string `json:"Status"`
		} `json:"Attachment,omitempty"`
	} `json:"NetworkInterfaces"`
}
