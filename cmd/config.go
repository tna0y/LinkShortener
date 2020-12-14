package main

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"

	"service/pkg/config"
)

func buildConfig() config.Config {
	viper.SetEnvPrefix("LS")
	_ = viper.BindEnv("BIND")
	_ = viper.BindEnv("BASE_URL")
	_ = viper.BindEnv("BOT_TOKEN")
	_ = viper.BindEnv("POSTGRES_DSN")

	return config.Config{
		Bind:          viper.GetString("BIND"),
		BaseURL:       viper.GetString("BASE_URL"),
		TelegramToken: viper.GetString("BOT_TOKEN"),
		PostgresDSN:   viper.GetString("POSTGRES_DSN"),
	}
}
