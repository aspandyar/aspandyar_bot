package util

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

// Config is the configuration for the application// It is used to configure the application
type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	val := reflect.ValueOf(config)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			err = fmt.Errorf("Missing required configuration: %s", val.Type().Field(i).Tag.Get("mapstructure"))
			return
		}
	}

	return
}
