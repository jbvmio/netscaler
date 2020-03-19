package netscaler

// NSAPIResponse represents the main portion of the Nitro API response
type NSAPIResponse struct {
	Errorcode                int64                     `json:"errorcode"`
	Message                  string                    `json:"message"`
	Severity                 string                    `json:"severity"`
	NSLicense                NSLicense                 `json:"nslicense"`
	NSStats                  NSStats                   `json:"ns"`
	InterfaceStats           []InterfaceStats          `json:"Interface"`
	VirtualServerStats       []VirtualServerStats      `json:"lbvserver"`
	ServiceStats             []ServiceStats            `json:"service"`
	ServiceGroups            []ServiceGroups           `json:"servicegroup"`
	ServiceGroupMemberStats  []ServiceGroupMemberStats `json:"servicegroupmember"`
	GSLBServiceStats         []GSLBServiceStats        `json:"gslbservice"`
	GSLBVirtualServerStats   []GSLBVirtualServerStats  `json:"gslbvserver"`
	CSVirtualServerStats     []CSVirtualServerStats    `json:"csvserver"`
	VPNVirtualServerStats    []VPNVirtualServerStats   `json:"vpnvserver"`
	LBVServerServiceBindings []LBVirtualServerBindings `json:"lbvserver_service_binding"`
}
