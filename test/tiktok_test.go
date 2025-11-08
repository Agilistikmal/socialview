package test

import (
	"testing"

	"github.com/agilistikmal/socialview/pkg/tiktok"
	"github.com/stretchr/testify/assert"
)

func TestTiktokGetVideo(t *testing.T) {
	tiktokService := tiktok.NewTiktokService()
	media, err := tiktokService.GetMedia("https://vt.tiktok.com/ZSyS5g8XL/")
	assert.NoError(t, err)
	assert.NotNil(t, media)

	t.Logf("Media ID: %s", media.ID)
}

func TestTiktokDownloadVideo(t *testing.T) {
	tiktokService := tiktok.NewTiktokService()
	err := tiktokService.Download("https://vt.tiktok.com/ZSyS5g8XL/", "./testdata/video.mp4")
	assert.NoError(t, err)
	assert.FileExists(t, "./testdata/video.mp4")
}

func TestTiktokGetImage(t *testing.T) {
	tiktokService := tiktok.NewTiktokService()
	media, err := tiktokService.GetMedia("https://vt.tiktok.com/ZSy9tWrUv/")
	assert.NoError(t, err)
	assert.NotNil(t, media)

	t.Logf("Media ID: %s", media.ID)
}
