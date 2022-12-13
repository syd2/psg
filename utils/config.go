package utils

import (
	"log"

	"github.com/spf13/viper"
)

//we will store the config from the env var in this struct
type Config struct {
	DBSource string `mapstructure:"DB_SOURCE"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	Secret   string `mapstructure:"SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("api")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cant read config")
		return
	}
	err = viper.Unmarshal(&config)
	return
}
