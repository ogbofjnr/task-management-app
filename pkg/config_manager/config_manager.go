package config_manager

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

// GetConfig render and return app configuration
func GetConfig(name string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(name)
	config.SetConfigType("toml")
	config.AddConfigPath("config")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}
