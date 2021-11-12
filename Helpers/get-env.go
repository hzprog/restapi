package helpers

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnvVar(key string) string {
	viper.SetConfigFile("config.env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
