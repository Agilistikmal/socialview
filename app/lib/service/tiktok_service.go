package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/agilistikmal/socialview/app/helper"
	"github.com/agilistikmal/socialview/app/model"
)

func GetTiktokVideoID(link string) string {
	r, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	u, _ := url.Parse(r.Request.URL.String())
	splitted := strings.Split(u.Path, "/")

	var id string
	for i, word := range splitted {
		if word == "video" {
			id = splitted[i+1]
		}
	}
	return id
}

func GetTiktokVideo(link string) model.Media {
	config := helper.LoadConfig()

	id := GetTiktokVideoID(link)
	apiUrl := config.SocialMedia["tiktok"].ApiUrl + id
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)

	var tiktokApiResponse model.TiktokAPIResponse
	json.Unmarshal(b, &tiktokApiResponse)

	media := model.Media{
		ID:       tiktokApiResponse.AwemeList[0].AwemeID,
		Type:     "video",
		Platform: "tiktok",
		Source:   tiktokApiResponse.AwemeList[0].Video.PlayAddr.UrlList[0],
	}
	return media
}
