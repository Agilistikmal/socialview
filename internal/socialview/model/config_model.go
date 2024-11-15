package model

type ConfigMedia struct {
	Name    string   `json:"name,omitempty"`
	Domains []string `json:"domains,omitempty"`
}
