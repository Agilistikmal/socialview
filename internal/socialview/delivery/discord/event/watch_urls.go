package event

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/agilistikmal/socialview/internal/socialview/model"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"mvdan.cc/xurls/v2"
)

func (h *DiscordEventHandler) WatchURLs(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}

	configMediaListJson, err := json.Marshal(viper.Get("media"))
	if err != nil {
		log.Fatal(err)
	}

	var configMediaList *[]model.ConfigMedia
	json.Unmarshal(configMediaListJson, &configMediaList)

	for _, mention := range msg.Mentions {
		if mention.ID == s.State.User.ID {
			rx := xurls.Relaxed()

			url := rx.FindString(msg.Content)

			if url == "" {
				s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
					Reference: msg.Reference(),
					Content:   "URL tidak ditemukan, sertakan link jika ingin mendownload",
				})
				return
			}

			for _, configMedia := range *configMediaList {
				for _, domain := range configMedia.Domains {
					if strings.Contains(url, domain) {
						s.ChannelTyping(msg.ID)

						msgWait, _ := s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
							Reference: msg.Reference(),
							Content:   fmt.Sprintf("Please wait! Downloading from %s", domain),
						})

						err := h.MediaService.DownloadMedia(configMedia.Name, url, "temp.mp4")
						if err != nil {
							s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
								Reference: msg.Reference(),
								Content:   fmt.Sprintf("Failed to download video: %s", err.Error()),
							})
							return
						}

						file, _ := os.Open("temp.mp4")

						s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
							Reference: msg.Reference(),
							File: &discordgo.File{
								Name:   file.Name(),
								Reader: file,
							},
						})

						s.ChannelMessageDelete(msgWait.ChannelID, msgWait.ID)
						os.Remove(file.Name())

						return
					}
				}
			}
			return
		}
	}

}
