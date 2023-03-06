package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var cf *Config

type VideoType = string

const (
	Movie VideoType = "movie"
)

type Config struct {
	DirPaths       []string  `json:"dir_paths"`
	SkipDataSize   int64     `json:"skip_data_size"`
	ScrapingWorker int64     `json:"scraping_worker"`
	VideoType      VideoType `json:"video_type"`
}

func Load() {
	cf = &Config{
		DirPaths:       nil,
		SkipDataSize:   0,
		ScrapingWorker: 10,
	}
	viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(cf, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(cf)
}

func Get() *Config {
	return cf
}
