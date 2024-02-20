package model

type InstagramAPIResponse struct {
	DisplayResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int    `json:"config_width"`
		ConfigHeight int    `json:"config_height"`
	} `json:"display_resources"`
	AccessibilityCaption string `json:"accessibility_caption"`
	IsVideo              bool   `json:"is_video"`
	VideoUrl             string `json:"video_url"`
}
