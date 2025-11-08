package tiktok

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type TiktokBaseRequestModel struct {
	WebIdLastTime   string `json:"WebIdLastTime,omitempty"`
	Aid             string `json:"aid,omitempty"`
	AppLanguage     string `json:"app_language,omitempty"`
	AppName         string `json:"app_name,omitempty"`
	BrowserLanguage string `json:"browser_language,omitempty"`
	BrowserName     string `json:"browser_name,omitempty"`
	BrowserOnline   string `json:"browser_online,omitempty"`
	BrowserPlatform string `json:"browser_platform,omitempty"`
	BrowserVersion  string `json:"browser_version,omitempty"`
	Channel         string `json:"channel,omitempty"`
	CookieEnabled   string `json:"cookie_enabled,omitempty"`
	DeviceId        int    `json:"device_id,omitempty"`
	OdinId          int    `json:"odinId,omitempty"`
	DevicePlatform  string `json:"device_platform,omitempty"`
	FocusState      string `json:"focus_state,omitempty"`
	FromPage        string `json:"from_page,omitempty"`
	HistoryLen      int    `json:"history_len,omitempty"`
	IsFullscreen    string `json:"is_fullscreen,omitempty"`
	IsPageVisible   string `json:"is_page_visible,omitempty"`
	Language        string `json:"language,omitempty"`
	PriorityRegion  string `json:"priority_region,omitempty"`
	Referer         string `json:"referer,omitempty"`
	Region          string `json:"region,omitempty"`
	RootReferer     string `json:"root_referer,omitempty"`
	ScreenHeight    int    `json:"screen_height,omitempty"`
	ScreenWidth     int    `json:"screen_width,omitempty"`
	WebcastLanguage string `json:"webcast_language,omitempty"`
	TzName          string `json:"tz_name,omitempty"`
	MsToken         string `json:"ms_token,omitempty"`
}

func (m *TiktokBaseRequestModel) Default(msToken string) *TiktokBaseRequestModel {
	return &TiktokBaseRequestModel{
		WebIdLastTime:   strconv.FormatInt(time.Now().Unix(), 10),
		Aid:             "1988",
		AppLanguage:     "en",
		AppName:         "tiktok_web",
		BrowserLanguage: "en-US",
		BrowserName:     "Mozilla",
		BrowserOnline:   "true",
		BrowserPlatform: "Win32",
		BrowserVersion:  "122.0.0.0",
		Channel:         "tiktok_web",
		CookieEnabled:   "true",
		DeviceId:        7380187414842836523,
		OdinId:          7404669909585003563,
		DevicePlatform:  "web_pc",
		FocusState:      "true",
		FromPage:        "user",
		HistoryLen:      4,
		IsFullscreen:    "false",
		IsPageVisible:   "true",
		Language:        "en",
		PriorityRegion:  "US",
		Referer:         "",
		Region:          "US",
		RootReferer:     "https://www.tiktok.com/",
		ScreenHeight:    1080,
		ScreenWidth:     1920,
		WebcastLanguage: "en",
		TzName:          "America/Tijuana",
		MsToken:         msToken,
	}
}

func (m *TiktokBaseRequestModel) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"WebIdLastTime":    m.WebIdLastTime,
		"aid":              m.Aid,
		"app_language":     m.AppLanguage,
		"app_name":         m.AppName,
		"browser_language": m.BrowserLanguage,
		"browser_name":     m.BrowserName,
		"browser_online":   m.BrowserOnline,
		"browser_platform": m.BrowserPlatform,
		"browser_version":  m.BrowserVersion,
		"channel":          m.Channel,
		"cookie_enabled":   m.CookieEnabled,
		"device_id":        m.DeviceId,
		"odinId":           m.OdinId,
		"device_platform":  m.DevicePlatform,
		"focus_state":      m.FocusState,
		"from_page":        m.FromPage,
		"history_len":      m.HistoryLen,
		"is_fullscreen":    m.IsFullscreen,
		"is_page_visible":  m.IsPageVisible,
		"language":         m.Language,
		"priority_region":  m.PriorityRegion,
		"referer":          m.Referer,
		"region":           m.Region,
		"root_referer":     m.RootReferer,
		"screen_height":    m.ScreenHeight,
		"screen_width":     m.ScreenWidth,
		"webcast_language": m.WebcastLanguage,
		"tz_name":          m.TzName,
		"ms_token":         m.MsToken,
	}
}

func (m *TiktokBaseRequestModel) ToQueryParams() string {
	params := url.Values{}
	for key, value := range m.ToMap() {
		params.Add(key, fmt.Sprintf("%v", value))
	}
	return params.Encode()
}
