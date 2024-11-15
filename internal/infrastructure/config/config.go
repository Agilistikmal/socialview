package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
