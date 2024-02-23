package model

type TiktokAPIResponse struct {
	AwemeList []struct {
		AwemeID string `json:"aweme_id"`
		Video   struct {
			PlayAddr struct {
				UrlList []string `json:"url_list"`
			} `json:"play_addr"`
		} `json:"video"`
		ImagePostInfo struct {
			Images []struct {
				DisplayImage struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"display_image"`
			} `json:"images"`
		} `json:"image_post_info"`
	} `json:"aweme_list"`
}
