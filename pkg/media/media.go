package media

import (
	"github.com/agilistikmal/socialview/pkg"
	"github.com/agilistikmal/socialview/pkg/tiktok"
)

type MediaService struct {
	tiktok pkg.SocialView
}

func NewMediaService() *MediaService {
	return &MediaService{
		tiktok: tiktok.NewTiktokService(),
	}
}
