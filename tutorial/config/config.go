package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type OpenSearchConfig struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type AppConfig struct {
	Server     ServerConfig     `mapstructure:"server"`
	OpenSearch OpenSearchConfig `mapstructure:"opensearch"`
}

// Global singleton-like config
var Cfg *AppConfig

func LoadConfig(env string) {
	v := viper.New()
	v.SetConfigName(env)
	v.SetConfigType("yaml")
	v.AddConfigPath("/app/resources")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var conf AppConfig
	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	Cfg = &conf
}
