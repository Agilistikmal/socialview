package pkg

type SocialView interface {
	GetMedia(url string) (*Media, error)
	Download(url string, filename string) error
}
