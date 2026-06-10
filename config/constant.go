package config

var ProtocolRepoMap = map[string]string{
	"nimiq": "https://github.com/beardsoft/nimiq-ansible.git",
	// Add other protocols here
}

var ProtocolNodeReleaseMap = map[string]string{
	"nimiq": "nimiq/core-rs-albatross",
	// Add other protocols here
}

var SupportedConfigurations = map[string]map[string][]string{
	"nimiq": {
		"testnet": {"validator", "full_node", "history_node"},
		"mainnet": {"validator", "full_node", "history_node"},
	},
	// Add other protocols here if needed
}
