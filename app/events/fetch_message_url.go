package events

import (
	"fmt"
	"os"
	"strings"

	"github.com/agilistikmal/socialview/app/helper"
	"github.com/agilistikmal/socialview/app/lib/service"
	"github.com/agilistikmal/socialview/app/model"
	"github.com/bwmarrin/discordgo"
	"mvdan.cc/xurls/v2"
)

func FetchMessageUrl(s *discordgo.Session, m *discordgo.MessageCreate) {
	config := helper.LoadConfig()

	if m.Author.Bot {
		return
	}

	var whitelist bool
	for _, ch := range config.AllowedChannels {
		if m.ChannelID == ch {
			whitelist = true
			break
		}
	}

	if !whitelist {
		return
	}
	if !strings.Contains(m.Message.Content, "http") {
		return
	}

	xurl := xurls.Strict()
	msgUrls := xurl.FindAllString(m.Message.Content, 1)

	for _, msgUrl := range msgUrls {
		for _, socialmedia := range config.SocialMedia {
			for _, url := range socialmedia.BaseUrl {
				if strings.Contains(msgUrl, url) {
					msg, _ := s.ChannelMessageSendReply(m.ChannelID, "Sedang mengambil video, tunggu sebentar...", &discordgo.MessageReference{
						MessageID: m.Message.ID,
						ChannelID: m.ChannelID,
						GuildID:   m.GuildID,
					})

					var media model.Media
					var filename string
					var file *os.File

					switch strings.ToLower(socialmedia.Name) {
					case "tiktok":
						media = service.GetTiktokVideo(msgUrl)
						filename = "./tmp/" + media.ID + ".mp4"
						service.SaveVideo(media.Source, filename)
						file = service.GetVideo(filename)
						defer file.Close()
					case "instagram":
						media = service.GetInstagramMedia(msgUrl)
						if media.Source == "ERROR" {
							content := m.Author.Mention() + " API Timeout Error, coba lagi."
							s.ChannelMessageEditComplex(&discordgo.MessageEdit{
								ID:      msg.ID,
								Channel: msg.ChannelID,
								Content: &content,
							})
							continue
						}
						if media.Type == "video" {
							filename = "./tmp/" + media.ID + ".mp4"
						} else {
							filename = "./tmp/" + media.ID + ".png"
						}
						service.SaveVideo(media.Source, filename)
						file = service.GetVideo(filename)
						defer file.Close()
					}

					embed := &discordgo.MessageEmbed{
						Description: fmt.Sprintf("From %s\n%s", m.Author.Mention(), m.Message.Content),
					}

					content := m.Author.Mention()
					s.ChannelMessageEditComplex(&discordgo.MessageEdit{
						ID:      msg.ID,
						Channel: msg.ChannelID,
						Content: &content,
						Files: []*discordgo.File{
							{
								Name:   filename,
								Reader: file,
							},
						},
						Embed: embed,
					})

					s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
					service.DeleteVideo(filename)
				}
			}
		}
	}
}
