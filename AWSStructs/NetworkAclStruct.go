package AWSStructs

type NetworkAcls struct {
	NetworkAcls []struct {
		Associations []struct {
			NetworkACLAssociationID string `json:"NetworkAclAssociationId"`
			NetworkACLID            string `json:"NetworkAclId"`
			SubnetID                string `json:"SubnetId"`
		} `json:"Associations"`
		Entries []struct {
			CidrBlock  string `json:"CidrBlock"`
			Egress     bool   `json:"Egress"`
			Protocol   string `json:"Protocol"`
			RuleAction string `json:"RuleAction"`
			RuleNumber int    `json:"RuleNumber"`
		} `json:"Entries"`
		IsDefault    bool   `json:"IsDefault"`
		NetworkACLID string `json:"NetworkAclId"`
		OwnerID      string `json:"OwnerId"`
		VpcID        string `json:"VpcId"`
	} `json:"NetworkAcls"`
}
