package config

// Raw structure of Configuration files
// log config files found
type Raw struct {
	Exec     string `json:"exec"`
	Commands struct {
		Autoremove string `json:"autoremove"`
		Build      string `json:"build"`
		Deps       string `json:"deps"`
		Depends    string `json:"depends"`
		Download   string `json:"download"`
		Fix        string `json:"fix"`
		Help       string `json:"help"`
		Info       string `json:"info"`
		Install    string `json:"install"`
		Installed  string `json:"installed"`
		Remove     string `json:"remove"`
		Search     string `json:"search"`
		Update     string `json:"update"`
		Upgrade    string `json:"upgrade"`
		Version    string `json:"version"`
	} `json:"commands"`
}
