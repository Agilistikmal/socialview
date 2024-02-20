package events

import (
	"fmt"
	"strings"

	"github.com/agilistikmal/socialview/app/helper"
	"github.com/agilistikmal/socialview/app/lib/service"
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
					video := service.GetTiktokVideo(msgUrl)
					filename := "./tmp/" + video.ID + ".mp4"
					service.SaveVideo(video.Source, filename)
					file := service.GetVideo(filename)
					defer file.Close()

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
								Name:        filename,
								ContentType: "video/mp4",
								Reader:      file,
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
