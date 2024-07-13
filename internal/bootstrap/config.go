package bootstrap

import (
	"github.com/spf13/viper"
	"go-boilerplate/internal/config"
	"log"
)

// NewConfig initialize config object
func NewConfig() *config.Config {

	cfg := config.Config{}

	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config/") // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		log.Fatalf("fatal error config file: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &cfg
}
