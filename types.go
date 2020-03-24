package netscaler

// ConfigType codes for Stats
type ConfigType int

// StatsType codes for Stats
type StatsType int

// ConfigType code list:
const (
	ConfigTypeNone ConfigType = iota
	ConfigTypeLicense
	ConfigTypeServiceGroup
	ConfigTypeLBVSBinding
	ConfigTypeLBVSSvcBinding
)

var configTypeStrings = [...]string{
	`configNone`,
	`nslicense`,
	`servicegroup`,
	`lbvserver_binding`,
	`lbvserver_service_binding`,
}

// StatsType code list:
const (
	StatsTypeNone StatsType = iota
	StatsTypeNS
	StatsTypeInterface
	StatsTypeLBVServer
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
