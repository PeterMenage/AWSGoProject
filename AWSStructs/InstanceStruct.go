package AWSStructs

type Instances struct {
	Instances []struct {
		AmiLaunchIndex      int    `json:"AmiLaunchIndex"`
		Architecture        string `json:"Architecture"`
		BlockDeviceMappings []struct {
			DeviceName string `json:"DeviceName"`
			Ebs        struct {
				AttachTime          string `json:"AttachTime"`
				DeleteOnTermination bool   `json:"DeleteOnTermination"`
				Status              string `json:"Status"`
				VolumeID            string `json:"VolumeId"`
			} `json:"Ebs"`
		} `json:"BlockDeviceMappings"`
		CapacityReservationSpecification struct {
			CapacityReservationPreference string `json:"CapacityReservationPreference"`
		} `json:"CapacityReservationSpecification"`
		ClientToken string `json:"ClientToken"`
		CPUOptions  struct {
			CoreCount      int `json:"CoreCount"`
			ThreadsPerCore int `json:"ThreadsPerCore"`
		} `json:"CpuOptions"`
		EbsOptimized       bool `json:"EbsOptimized"`
		EnaSupport         bool `json:"EnaSupport"`
		HibernationOptions struct {
			Configured bool `json:"Configured"`
		} `json:"HibernationOptions"`
		Hypervisor   string `json:"Hypervisor"`
		ImageID      string `json:"ImageId"`
		InstanceID   string `json:"InstanceId"`
		InstanceType string `json:"InstanceType"`
		KeyName      string `json:"KeyName"`
		LaunchTime   string `json:"LaunchTime"`
		Monitoring   struct {
			State string `json:"State"`
		} `json:"Monitoring"`
		NetworkInterfaces []struct {
			Association struct {
				IPOwnerID     string `json:"IpOwnerId"`
				PublicDNSName string `json:"PublicDnsName"`
				PublicIP      string `json:"PublicIp"`
			} `json:"Association"`
			Attachment struct {
				AttachTime          string `json:"AttachTime"`
				AttachmentID        string `json:"AttachmentId"`
				DeleteOnTermination bool   `json:"DeleteOnTermination"`
				DeviceIndex         int    `json:"DeviceIndex"`
				Status              string `json:"Status"`
			} `json:"Attachment"`
			Description string `json:"Description"`
			Groups      []struct {
				GroupID   string `json:"GroupId"`
				GroupName string `json:"GroupName"`
			} `json:"Groups"`
			MacAddress         string `json:"MacAddress"`
			NetworkInterfaceID string `json:"NetworkInterfaceId"`
			OwnerID            string `json:"OwnerId"`
			PrivateDNSName     string `json:"PrivateDnsName"`
			PrivateIPAddress   string `json:"PrivateIpAddress"`
			PrivateIPAddresses []struct {
				Association struct {
					IPOwnerID     string `json:"IpOwnerId"`
					PublicDNSName string `json:"PublicDnsName"`
					PublicIP      string `json:"PublicIp"`
				} `json:"Association"`
				Primary          bool   `json:"Primary"`
				PrivateDNSName   string `json:"PrivateDnsName"`
				PrivateIPAddress string `json:"PrivateIpAddress"`
			} `json:"PrivateIpAddresses"`
			SourceDestCheck bool   `json:"SourceDestCheck"`
			Status          string `json:"Status"`
			SubnetID        string `json:"SubnetId"`
			VpcID           string `json:"VpcId"`
		} `json:"NetworkInterfaces"`
		Placement struct {
			AvailabilityZone string `json:"AvailabilityZone"`
			GroupName        string `json:"GroupName"`
			Tenancy          string `json:"Tenancy"`
		} `json:"Placement"`
		PrivateDNSName   string `json:"PrivateDnsName"`
		PrivateIPAddress string `json:"PrivateIpAddress"`
		PublicDNSName    string `json:"PublicDnsName"`
		PublicIPAddress  string `json:"PublicIpAddress"`
		RootDeviceName   string `json:"RootDeviceName"`
		RootDeviceType   string `json:"RootDeviceType"`
		SecurityGroups   []struct {
			GroupID   string `json:"GroupId"`
			GroupName string `json:"GroupName"`
		} `json:"SecurityGroups"`
		SourceDestCheck bool `json:"SourceDestCheck"`
		State           struct {
			Code int    `json:"Code"`
			Name string `json:"Name"`
		} `json:"State"`
		StateTransitionReason string `json:"StateTransitionReason"`
		SubnetID              string `json:"SubnetId"`
		Tags                  []struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
		} `json:"Tags"`
		VirtualizationType string `json:"VirtualizationType"`
		VpcID              string `json:"VpcId"`
	} `json:"Instances"`
	OwnerID       string `json:"OwnerId"`
	ReservationID string `json:"ReservationId"`
}
