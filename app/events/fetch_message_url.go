package events

import (
	"strings"

	"github.com/agilistikmal/socialview/app/helper"
	"github.com/agilistikmal/socialview/app/lib/service"
	"github.com/bwmarrin/discordgo"
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

	splitted := strings.Split(m.Message.Content, " ")
	for _, word := range splitted {
		for _, socialmedia := range config.SocialMedia {
			for _, url := range socialmedia.BaseUrl {
				if strings.Contains(word, url) {
					msg, _ := s.ChannelMessageSendReply(m.ChannelID, "Sedang mengambil video, tunggu sebentar...", &discordgo.MessageReference{
						MessageID: m.Message.ID,
						ChannelID: m.ChannelID,
						GuildID:   m.GuildID,
					})
					tiktokVideo := service.GetTiktokVideo(word)
					videoUrl := tiktokVideo.AwemeList[0].Video.PlayAddr.UrlList[0]

					filename := "./tmp/video.mp4"
					service.SaveVideo(videoUrl, filename)
					file := service.GetVideo(filename)
					defer file.Close()

					content := tiktokVideo.AwemeList[0].AwemeID
					s.ChannelMessageEditComplex(&discordgo.MessageEdit{
						ID:      msg.ID,
						Channel: msg.ChannelID,
						Content: &content,
						Files: []*discordgo.File{
							{
								Name:        "tiktok.mp4",
								ContentType: "video/mp4",
								Reader:      file,
							},
						},
					})
					service.DeleteVideo(filename)
				}
			}
		}
	}
}
