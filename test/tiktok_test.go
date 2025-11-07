package test

import (
	"testing"

	"github.com/agilistikmal/socialview/pkg/tiktok"
	"github.com/stretchr/testify/assert"
)

func TestTiktokGetMedia(t *testing.T) {
	tiktokService := tiktok.NewTiktokService()
	media, err := tiktokService.GetMedia("https://vt.tiktok.com/ZSyS5g8XL/")
	assert.NoError(t, err)
	assert.NotNil(t, media)

	t.Logf("Media ID: %s", media.ID)
}

func TestTiktokDownload(t *testing.T) {
	tiktokService := tiktok.NewTiktokService()
	err := tiktokService.Download("https://vt.tiktok.com/ZSyS5g8XL/", "./testdata/video.mp4")
	assert.NoError(t, err)
	assert.FileExists(t, "./testdata/video.mp4")
}
