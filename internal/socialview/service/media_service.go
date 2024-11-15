package service

import "fmt"

type MediaService struct {
	TikTokService *TikTokService
}

func NewMediaService(tiktokService *TikTokService) *MediaService {
	return &MediaService{
		TikTokService: tiktokService,
	}
}

func (s *MediaService) DownloadMedia(media string, url string, filename string) error {
	switch media {
	case "tiktok":
		err := s.TikTokService.Download(url, filename)
		return err
	default:
		return fmt.Errorf("invalid media")
	}
}
