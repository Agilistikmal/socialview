package helper

import (
	"encoding/json"
	"io"
	"os"

	"github.com/agilistikmal/socialview/app/model"
)

func LoadConfig() *model.Config {
	var config model.Config

	f, _ := os.Open("./config.json")
	defer f.Close()

	b, _ := io.ReadAll(f)
	json.Unmarshal(b, &config)

	return &config
}
