package pkg

type Platform string

const (
	PlatformTiktok Platform = "tiktok"
)

type MediaType string

const (
	MediaTypeVideo    MediaType = "video"
	MediaTypeImage    MediaType = "image"
	MediaTypeAudio    MediaType = "audio"
	MediaTypeDocument MediaType = "document"
	MediaTypeOther    MediaType = "other"
)

type Media struct {
	Platform     Platform    `json:"platform,omitempty"`
	MediaTypes   []MediaType `json:"media_types,omitempty"`
	URL          string      `json:"url,omitempty"`
	ID           string      `json:"id,omitempty"`
	Title        string      `json:"title,omitempty"`
	Description  string      `json:"description,omitempty"`
	ThumbnailURL string      `json:"thumbnail_url,omitempty"`
	MediaItems   []MediaItem `json:"media_items,omitempty"`
	Duration     int         `json:"duration,omitempty"`
	Views        int         `json:"views,omitempty"`
	Likes        int         `json:"likes,omitempty"`
	Comments     int         `json:"comments,omitempty"`
	Author       Author      `json:"author,omitempty"`
}

type MediaItem struct {
	URL  string    `json:"url,omitempty"`
	ID   string    `json:"id,omitempty"`
	Type MediaType `json:"type,omitempty"`
}

type Author struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	URL       string `json:"url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}
