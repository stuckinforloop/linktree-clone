package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	DB   string `mapstructure:"DB_TYPE"`
	MURL string `mapstructure:"MONGO_URL"`
	RURL string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
		return
	}
	err = viper.Unmarshal(&config)
	return
}
