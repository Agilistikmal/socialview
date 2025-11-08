package tiktok

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	cookies, err := s.getCookies(u)
	if err != nil {
		return nil, err
	}
	jar, _ := cookiejar.New(nil)
	for _, cookie := range cookies {
		jar.SetCookies(parsedUrl, []*http.Cookie{cookie})
	}
	s.httpClient.Jar = jar

	postId, _, err := s.getPostIdAndType(u)
	if err != nil {
		return nil, err
	}

	msToken := ""
	for _, cookie := range cookies {
		if cookie.Name == "msToken" {
			msToken = cookie.Value
			break
		}
	}

	baseRequestModel := &TiktokBaseRequestModel{}
	baseRequestModel = baseRequestModel.Default(msToken)
	params := baseRequestModel.ToQueryParams() + "&itemId=" + postId

	xbogusParams, _, _ := s.xbogusService.GetXBogus("/api/item/detail?" + params)

	endpoint := fmt.Sprintf("%s%s", s.baseUrl, xbogusParams)

	log.Printf("Endpoint: %s", endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Referer", s.referer)
	req.Header.Set("User-Agent", s.userAgent)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Response Body: %s", string(responseBody))

	return nil, nil
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
