package event

import "github.com/agilistikmal/socialview/internal/socialview/service"

type DiscordEventHandler struct {
	MediaService *service.MediaService
}

func NewDiscordEventHandler(mediaService *service.MediaService) *DiscordEventHandler {
	return &DiscordEventHandler{
		MediaService: mediaService,
	}
}
