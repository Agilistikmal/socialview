package tiktok

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"

	"github.com/agilistikmal/socialview/pkg"
	"github.com/agilistikmal/socialview/pkg/xbogus"
	"github.com/gocolly/colly/v2"
)

type TiktokService struct {
	collector     *colly.Collector
	baseUrl       string
	uaKey         string
	userAgent     string
	referer       string
	httpClient    *http.Client
	xbogusService *xbogus.XBogusService
}

func NewTiktokService() pkg.SocialView {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"
	return &TiktokService{
		collector:     colly.NewCollector(),
		baseUrl:       "https://www.tiktok.com",
		uaKey:         "df6b28f9b5f8a8e8c4c4c4c4c4c4c4c4",
		userAgent:     userAgent,
		referer:       "https://www.tiktok.com/",
		httpClient:    &http.Client{},
		xbogusService: xbogus.NewXBogusService(userAgent),
	}
}

func (s *TiktokService) GetMedia(u string) (*pkg.Media, error) {
	var media *pkg.Media
	s.collector.OnHTML("script#__UNIVERSAL_DATA_FOR_REHYDRATION__", func(h *colly.HTMLElement) {
		var universalData TiktokUniversalData

		err := json.Unmarshal([]byte(h.Text), &universalData)
		if err != nil {
			log.Fatal(err)
		}

		video := universalData.DefaultScope.VideoDetail.ItemInfo.ItemStruct.Video
		media = video.ToMedia()
	})

	s.collector.Visit(u)

	return media, nil
}

func (s *TiktokService) getCookies(u string) ([]*http.Cookie, error) {
	resp, err := s.httpClient.Head(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp.Cookies(), nil
}

func (s *TiktokService) getPostIdAndType(u string) (string, TiktokPostType, error) {
	resp, err := http.Head(u)
	if err != nil {
		return "", TiktokPostTypeOther, err
	}
	defer resp.Body.Close()

	path := resp.Request.URL.Path
	parts := strings.Split(path, "/")
	postType := TiktokPostType(parts[len(parts)-2])
	postId := parts[len(parts)-1]

	return postId, postType, nil
}

func (s *TiktokService) Download(u string, filename string) error {
	videoCollector := s.collector.Clone()

	s.collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", "https://www.tiktok.com/")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
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

	s.collector.Visit(u)

	videoCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", u)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	})

	videoCollector.OnResponse(func(r *colly.Response) {
		os.MkdirAll(filepath.Dir(filename), 0755)
		os.WriteFile(filename, r.Body, 0644)
	})

	return nil
}
