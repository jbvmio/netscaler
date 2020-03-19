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
)

var configTypeStrings = [...]string{
	`ConfigNone`,
	`nslicense`,
	`servicegroup`,
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
)

var statsTypeStrings = [...]string{
	`StatsNone`,
	`ns`,
	`interface`,
}

func (t ConfigType) String() string {
	return configTypeStrings[t]
}

func (t StatsType) String() string {
	return statsTypeStrings[t]
}
