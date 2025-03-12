package bootstrap

import (
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"go-boilerplate/internal/config"

	"github.com/spf13/viper"
)

var configData *config.Config

// NewConfig initialize config object
func NewConfig() *config.Config {

	viper.SetConfigType("env")
	viper.SetConfigName(".env") // name of Config file (without extension)
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // if env keys are set in the system, use it instead of config file

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Warnf("error reading config file: %v", err)
	}

	cfg := config.Config{
		App: config.App{
			SecretKey:    getRequiredString("APP_SECRET_KEY"),
			Env:          GetEnv("APP_ENV", "local"),
			Name:         GetEnv("APP_NAME", "go-boilerplate"),
			Port:         GetEnv("APP_PORT", "8080"),
			ReadTimeout:  getEnvAsInt("APP_READ_TIMEOUT", 10),
			WriteTimeout: getEnvAsInt("APP_WRITE_TIMEOUT", 10),
			ReleaseMode:  GetEnv("APP_RELEASE_MODE", "debug"),
		},
		Postgre: config.Postgre{
			IsEnabled:   getRequiredBool("POSTGRE_IS_ENABLED"),
			URL:         getRequiredString("POSTGRE_URL"),
			MaxIdleConn: getEnvAsInt("POSTGRE_MAX_IDLE_CONN", 10),
			MaxOpenConn: getEnvAsInt("POSTGRE_MAX_OPEN_CONN", 10),
		},
		JWT: config.JWT{
			SecretKey: getRequiredString("JWT_SECRET_KEY"),
		},
	}
	return &cfg
}

func getRequiredString(key string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}

	log.Fatalln(fmt.Errorf("KEY %s IS MISSING", key))
	return ""
}

func getRequiredInt(key string) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}

func getRequiredTime(key string) time.Duration {
	if viper.IsSet(key) {
		return time.Duration(viper.GetInt(key)) * time.Second
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}

func getRequiredBool(key string) bool {
	if viper.IsSet(key) {
		return viper.GetBool(key)
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return value
	}

	return defaultVal
}

func GetConfig() *config.Config {
	return configData
}
