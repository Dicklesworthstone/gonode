package pastel

// MasterNodeConfig represents pastel masternode configuration.
type MasterNodeConfig struct {
	Alias       string `json:"alias"`
	Address     string `json:"address"`
	PrivateKey  string `json:"privateKey"`
	TXHash      string `json:"txHash"`
	OutputIndex string `json:"outputIndex"`
	ExtAddress  string `json:"extAddress"`
	ExtKey      string `json:"extKey"`
	ExtP2P      string `json:"extP2P"`
	ExtCfg      string `json:"extCfg"`
	Status      string `json:"status"`
}
