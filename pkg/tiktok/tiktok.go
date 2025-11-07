package tiktok

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"

	"github.com/agilistikmal/socialview/pkg"
	"github.com/gocolly/colly/v2"
)

type TiktokService struct {
	collector *colly.Collector
}

func NewTiktokService() pkg.SocialView {
	return &TiktokService{
		collector: colly.NewCollector(),
	}
}

func (s *TiktokService) GetMedia(url string) (*pkg.Media, error) {
	var video *TiktokVideo
	s.collector.OnHTML("script#__UNIVERSAL_DATA_FOR_REHYDRATION__", func(h *colly.HTMLElement) {
		var universalData TiktokUniversalData

		err := json.Unmarshal([]byte(h.Text), &universalData)
		if err != nil {
			log.Fatal(err)
		}

		video = &universalData.DefaultScope.VideoDetail.ItemInfo.ItemStruct.Video
	})

	s.collector.Visit(url)

	return video.ToMedia(), nil
}

func (s *TiktokService) Download(url string, filename string) error {
	videoCollector := s.collector.Clone()

	videoCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", url)
	})

	videoCollector.OnResponse(func(r *colly.Response) {
		os.MkdirAll(filepath.Dir(filename), 0755)
		os.WriteFile(filename, r.Body, 0644)
	})

	s.collector.OnHTML("script#__UNIVERSAL_DATA_FOR_REHYDRATION__", func(h *colly.HTMLElement) {
		var universalData TiktokUniversalData

		err := json.Unmarshal([]byte(h.Text), &universalData)
		if err != nil {
			log.Fatal(err)
		}

		video := universalData.DefaultScope.VideoDetail.ItemInfo.ItemStruct.Video

		setCookiesHeaders := h.Response.Headers.Values("Set-Cookie")
		setCookies := make([]*http.Cookie, 0)
		for _, setCookieHeader := range setCookiesHeaders {
			cookies, err := http.ParseSetCookie(setCookieHeader)
			if err != nil {
				log.Fatalf("%s\nfailed to parse cookie: %v", setCookieHeader, err)
			}
			setCookies = append(setCookies, cookies)
		}
		videoCollector.SetCookies(h.Request.URL.String(), setCookies)

		videoCollector.Visit(video.PlayAddr)
	})

	s.collector.Visit(url)

	return nil
}
