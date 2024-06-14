package config

func IsValidConfiguration(protocol, network, nodeType string) bool {
	networks, ok := SupportedConfigurations[protocol]
	if !ok {
		return false
	}
	nodeTypes, ok := networks[network]
	if !ok {
		return false
	}
	for _, nt := range nodeTypes {
		if nt == nodeType {
			return true
		}
	}
	return false
}
