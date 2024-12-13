package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	MSite mSiteConfig `yaml:"m_site"`
}

type mSiteConfig struct {
	Title         string `yaml:"title"`
	Author        string `yaml:"author"`
	RecentPostNum int    `yaml:"recent_post_num"`
}

var MConfigInstance *MConfig

func init() {
	configFile, err := os.OpenFile("config.yaml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	configByte, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configByte, &MConfigInstance)
	if err != nil {
		log.Fatal(err)
	}
}
