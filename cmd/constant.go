package cmd

var ProtocolRepoMap = map[string]string{
	"nimiq": "https://github.com/Beardsoft/nimiq-ansible.git",
	// Add other protocols here
}

var SupportedConfigurations = map[string]map[string][]string{
	"nimiq": {
		"testnet": {"validator", "full_node", "history_node"},
		// Add other networks and node types here if needed
	},
	// Add other protocols here if needed
}
