package AWSStructs

type RouteTable struct {
	RouteTables []struct {
		Associations []struct {
			Main                    bool   `json:"Main"`
			RouteTableAssociationID string `json:"RouteTableAssociationId"`
			RouteTableID            string `json:"RouteTableId"`
		} `json:"Associations"`
		OwnerID      string `json:"OwnerId"`
		RouteTableID string `json:"RouteTableId"`
		Routes       []struct {
			DestinationCidrBlock string `json:"DestinationCidrBlock"`
			GatewayID            string `json:"GatewayId"`
			Origin               string `json:"Origin"`
			State                string `json:"State"`
		} `json:"Routes"`
		VpcID string `json:"VpcId"`
	} `json:"RouteTables"`
}
