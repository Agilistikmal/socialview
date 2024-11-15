package test

import (
	"log"
	"testing"

	"github.com/agilistikmal/socialview/internal/socialview/service"
)

func TestDownload(t *testing.T) {
	s := service.NewTikTokService()
	err := s.Download("https://vt.tiktok.com/ZSjf3jFm3/", "video.mp4")
	if err != nil {
		log.Fatal(err)
	}
}
