package netscaler

// ConfigType codes for Stats
type ConfigType int

// StatsType codes for Stats
type StatsType int

// ConfigType code list:
const (
	ConfigTypeNone ConfigType = iota
	ConfigTypeLicense
	ConfigTypeNSVersion
	ConfigTypeService
	ConfigTypeServiceGroup
	ConfigTypeLBVSBinding
	ConfigTypeLBVSSvcBinding
	ConfigTypeNSHardware
)

var configTypeStrings = [...]string{
	`configNone`,
	`nslicense`,
	`nsversion`,
	`service`,
	`servicegroup`,
	`lbvserver_binding`,
	`lbvserver_service_binding`,
	`nshardware`,
}

// StatsType code list:
const (
	StatsTypeNone StatsType = iota
	StatsTypeNS
	StatsTypeInterface
	StatsTypeLBVServer
	StatsTypeLBVServerWithStatBindings
	StatsTypeService
	StatsTypeServiceGroupMember
	StatsTypeGSLBService
	StatsTypeGSLBVServer
	StatsTypeCSVServer
	StatsTypeVPNVServer
	StatsTypeSSL
)

var statsTypeStrings = [...]string{
	`statsNone`,
	`ns`,
	`interface`,
	`lbvserver`,
	`lbvserver`,
	`service`,
	`servicegroup`,
	`gslbservice`,
	`gslbvserver`,
	`csvserver`,
	`vpnvserver`,
	`ssl`,
}

func (t ConfigType) String() string {
	return configTypeStrings[t]
}

func (t StatsType) String() string {
	return statsTypeStrings[t]
}

// NitroType represents either a StatsType or ConfigType.
type NitroType interface {
	String() string
}

// makeURL constructs a URL based on the given NitroType.
func (c *NitroClient) makeURL(nitroType NitroType) string {
	switch nitroType.(type) {
	case StatsType:
		return c.url + "stat/" + nitroType.String()
	case ConfigType:
		return c.url + "config/" + nitroType.String()
	default:
		return nitroType.String()
	}
}
