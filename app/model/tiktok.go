package model

type TiktokAPIResponse struct {
	AwemeList []struct {
		AwemeID string `json:"aweme_id"`
		Video   struct {
			PlayAddr struct {
				UrlList []string `json:"url_list"`
			} `json:"play_addr"`
		} `json:"video"`
	} `json:"aweme_list"`
}
