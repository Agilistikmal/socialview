package tiktok

import (
	"github.com/agilistikmal/socialview/pkg"
)

type TiktokPostType string

const (
	TiktokPostTypeVideo    TiktokPostType = "video"
	TiktokPostTypePhoto    TiktokPostType = "photo"
	TiktokPostTypeAudio    TiktokPostType = "audio"
	TiktokPostTypeDocument TiktokPostType = "document"
	TiktokPostTypeOther    TiktokPostType = "other"
)

type TiktokUniversalData struct {
	DefaultScope TiktokDefaultScope `json:"__DEFAULT_SCOPE__,omitempty"`
}

type TiktokDefaultScope struct {
	VideoDetail TiktokVideoDetail `json:"webapp.video-detail,omitempty"`
}

type TiktokVideoDetail struct {
	ItemInfo TiktokItemInfo `json:"itemInfo,omitempty"`
}

type TiktokItemInfo struct {
	ItemStruct TiktokItemStruct `json:"itemStruct,omitempty"`
}

type TiktokItemStruct struct {
	ID           string      `json:"id,omitempty"`
	Desc         string      `json:"desc,omitempty"`
	CreateTime   string      `json:"createTime,omitempty"`
	ScheduleTime int         `json:"scheduleTime,omitempty"`
	Video        TiktokVideo `json:"video,omitempty"`
}

type TiktokVideo struct {
	ID           string `json:"id,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
	Duration     int    `json:"duration,omitempty"`
	Ratio        string `json:"ratio,omitempty"`
	Cover        string `json:"cover,omitempty"`
	OriginCover  string `json:"originCover,omitempty"`
	DynamicCover string `json:"dynamicCover,omitempty"`
	PlayAddr     string `json:"playAddr,omitempty"`
	Bitrate      int    `json:"bitrate,omitempty"`
	Format       string `json:"format,omitempty"`
	VideoQuality string `json:"videoQuality,omitempty"`
	CodecType    string `json:"codecType,omitempty"`
	Definition   string `json:"definition,omitempty"`
}

func (v *TiktokVideo) ToMedia() *pkg.Media {
	return &pkg.Media{
		Platform:     pkg.PlatformTiktok,
		MediaTypes:   []pkg.MediaType{pkg.MediaTypeVideo},
		URL:          v.PlayAddr,
		ID:           v.ID,
		Title:        "",
		Description:  "",
		ThumbnailURL: v.Cover,
		Duration:     v.Duration,
		Views:        0,
		Likes:        0,
		Comments:     0,
		Author: pkg.Author{
			ID:        "",
			Username:  "",
			FullName:  "",
			URL:       "",
			AvatarURL: "",
		},
		MediaItems: []pkg.MediaItem{
			{
				URL:  v.PlayAddr,
				ID:   v.ID,
				Type: pkg.MediaTypeVideo,
			},
		},
	}
}
