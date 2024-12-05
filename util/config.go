package util

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	TelegramToken string `mapstructure:"TELEGRAM_TOKEN"`
	OpenaiToken   string `mapstructure:"OPENAI_TOKEN"`
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
			err = fmt.Errorf("missing required configuration: %s", val.Type().Field(i).Tag.Get("mapstructure"))
			return
		}
	}

	return
}

func LoadPromptByName(name string) (string, error) {
	filePath := fmt.Sprintf("./notes/%s.txt", name)

	files, err := os.ReadDir(".")
	if err != nil {
		return "", fmt.Errorf("cannot read directory: %v", err)
	}

	fmt.Println("Files in ./notes directory:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	prompt, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot read prompt: %v", err)
	}

	return string(prompt), nil
}
