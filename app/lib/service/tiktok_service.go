package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/agilistikmal/socialview/app/helper"
	"github.com/agilistikmal/socialview/app/model"
)

func GetTiktokVideoID(url string) string {
	var id string
	splitted := strings.Split(url, "/")
	for i, word := range splitted {
		if word == "video" {
			id = splitted[i+1]
		}
	}
	return id
}

func GetTiktokVideo(url string) *model.TiktokAPIResponse {
	config := helper.LoadConfig()

	id := GetTiktokVideoID(url)
	apiUrl := config.SocialMedia["tiktok"].ApiUrl + id
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)

	var tiktokApiResponse model.TiktokAPIResponse
	json.Unmarshal(b, &tiktokApiResponse)

	return &tiktokApiResponse
}
