package ports

// ConfigPort is the configure for MAO.
type ConfigPort interface {
	FormatConfig(corrConf *CorrelatorConfig, swConf *SwitchConfig) (*WFSConfig, error)
}

// WFSConfig represents the configure for MAO.
type WFSConfig struct {
	Correlator CorrelatorConfig
	Switch     SwitchConfig
}

// CorrelatorConfig represents the configure for the correlator of MAO.
type CorrelatorConfig struct {
	IntegrationTime IntegrationTime
	WindowFunction  WindowFunction
}

// IntegrationTime is the integration time in a correlation.
type IntegrationTime uint8

// WindowFunction is the window function for a FFT in a correlation.
type WindowFunction string

// SwitchConfig represents the configure for the switch of MAO.
type SwitchConfig struct {
	SwitchOrder SwitchOrder
}

// SwitchOrder is the order to switch.
type SwitchOrder []string
