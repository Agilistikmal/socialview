package model

type TikTokUniversalData struct {
	DefaultScope TikTokDefaultScope `json:"__DEFAULT_SCOPE__,omitempty"`
}

type TikTokDefaultScope struct {
	VideoDetail TikTokVideoDetail `json:"webapp.video-detail,omitempty"`
}

type TikTokVideoDetail struct {
	ItemInfo TikTokItemInfo `json:"itemInfo,omitempty"`
}

type TikTokItemInfo struct {
	ItemStruct TikTokItemStruct `json:"itemStruct,omitempty"`
}

type TikTokItemStruct struct {
	ID           string      `json:"id,omitempty"`
	Desc         string      `json:"desc,omitempty"`
	CreateTime   string      `json:"createTime,omitempty"`
	ScheduleTime int         `json:"scheduleTime,omitempty"`
	Video        TikTokVideo `json:"video,omitempty"`
}

type TikTokVideo struct {
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
