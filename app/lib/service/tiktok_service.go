package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/agilistikmal/socialview/app/helper"
	"github.com/agilistikmal/socialview/app/model"
	"github.com/chromedp/chromedp"
)

func GetTiktokMediaID(link string) (string, string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var location string
	task := chromedp.Tasks{
		chromedp.Navigate(link),
		chromedp.WaitVisible("#main-content-video_detail", chromedp.ByQuery),
		chromedp.Location(&location),
	}
	chromedp.Run(ctx, task)

	u, _ := url.Parse(location)
	splitted := strings.Split(u.Path, "/")
	id := splitted[3]

	var mediaType string
	switch splitted[2] {
	case "video":
		mediaType = "video"
	case "story":
		mediaType = "video"
	case "photo":
		mediaType = "image"
	}

	return mediaType, id
}

func GetTiktokMedia(link string) []model.Media {
	config := helper.LoadConfig()

	mediaType, id := GetTiktokMediaID(link)
	apiUrl := config.SocialMedia["tiktok"].ApiUrl + id
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)

	var tiktokApiResponse model.TiktokAPIResponse
	json.Unmarshal(b, &tiktokApiResponse)

	tiktokMedia := tiktokApiResponse.AwemeList[0]
	if tiktokMedia.AwemeID != id {
		return nil
	}

	var mediaList []model.Media

	if mediaType == "video" {
		media := model.Media{
			ID:       tiktokMedia.AwemeID,
			Type:     mediaType,
			Platform: "tiktok",
			Source:   tiktokMedia.Video.PlayAddr.UrlList[0],
		}
		mediaList = append(mediaList, media)
	} else if mediaType == "image" {
		for _, image := range tiktokMedia.ImagePostInfo.Images {
			mediaID := strings.Split(image.DisplayImage.Uri, "/")[1]
			media := model.Media{
				ID:       mediaID,
				Type:     mediaType,
				Platform: "tiktok",
				Source:   image.DisplayImage.UrlList[0],
			}
			mediaList = append(mediaList, media)
		}
	}

	return mediaList
}
