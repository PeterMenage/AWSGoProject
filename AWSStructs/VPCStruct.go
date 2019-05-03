package AWSStructs

type VPC struct {
	Vpcs []struct {
		CidrBlock               string `json:"CidrBlock"`
		CidrBlockAssociationSet []struct {
			AssociationID  string `json:"AssociationId"`
			CidrBlock      string `json:"CidrBlock"`
			CidrBlockState struct {
				State string `json:"State"`
			} `json:"CidrBlockState"`
		} `json:"CidrBlockAssociationSet"`
		DhcpOptionsID   string `json:"DhcpOptionsId"`
		InstanceTenancy string `json:"InstanceTenancy"`
		IsDefault       bool   `json:"IsDefault"`
		OwnerID         string `json:"OwnerId"`
		State           string `json:"State"`
		VpcID           string `json:"VpcId"`
	} `json:"Vpcs"`
}
