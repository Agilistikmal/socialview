package model

type Config struct {
	Bot             Bot                    `json:"bot"`
	SocialMedia     map[string]SocialMedia `json:"social_media"`
	AllowedChannels []string               `json:"allowed_channels"`
}

type Bot struct {
	Name         string `json:"name"`
	CustomStatus string `json:"custom_status"`
}

type SocialMedia struct {
	Name    string   `json:"name"`
	BaseUrl []string `json:"base_url"`
	ApiUrl  string   `json:"api_url"`
}
