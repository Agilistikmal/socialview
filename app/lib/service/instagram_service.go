package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/agilistikmal/socialview/app/helper"
	"github.com/agilistikmal/socialview/app/model"
	"github.com/joho/godotenv"
)

func GetInstagramMedia(link string) []model.Media {
	godotenv.Load()
	config := helper.LoadConfig()

	apiUrl := config.SocialMedia["instagram"].ApiUrl + link

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("X-RapidAPI-Key", os.Getenv("IG_RAPID_KEY"))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)

	var instagramResponse model.InstagramAPIResponse
	err = json.Unmarshal(b, &instagramResponse)
	if err != nil {
		log.Fatal(err)
	}

	if len(instagramResponse.DisplayResources) == 0 {
		return nil
	}

	var mediaType string
	var source string
	if instagramResponse.IsVideo {
		mediaType = "video"
		source = instagramResponse.VideoUrl
	} else {
		mediaType = "image"
		source = instagramResponse.DisplayResources[len(instagramResponse.DisplayResources)-1].Src
	}

	var mediaList []model.Media
	media := model.Media{
		Type:     mediaType,
		Platform: "instagram",
		Source:   source,
	}
	mediaList = append(mediaList, media)

	return mediaList
}
