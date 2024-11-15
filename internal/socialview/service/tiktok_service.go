package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/agilistikmal/socialview/internal/socialview/model"
	"github.com/gocolly/colly/v2"
)

type TikTokService struct {
}

func NewTikTokService() *TikTokService {
	return &TikTokService{}
}

func (s *TikTokService) Download(url string, filename string) error {
	collector := colly.NewCollector()
	videoCollector := collector.Clone()

	collector.OnHTML("script#__UNIVERSAL_DATA_FOR_REHYDRATION__", func(h *colly.HTMLElement) {
		var universalData model.TikTokUniversalData

		err := json.Unmarshal([]byte(h.Text), &universalData)
		if err != nil {
			log.Fatal(err)
		}

		video := universalData.DefaultScope.VideoDetail.ItemInfo.ItemStruct.Video

		videoCollector.Visit(video.PlayAddr)
	})

	videoCollector.OnResponse(func(r *colly.Response) {
		os.WriteFile(filename, r.Body, 0644)
	})

	collector.Visit(url)

	return nil
}
