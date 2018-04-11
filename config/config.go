package config

import (
	"log"

	"github.com/spf13/viper"
)

var item Configuration

// Configuration contains all the config content
type (
	Configuration struct {
		Mysql mysqlConfig
		Redis redisConfig
		Jwt   jwtConfig
	}

	mysqlConfig struct {
		Host            string
		Port            int
		Dbname          string
		Username        string
		Password        string
		MaxIdelConns    int
		MaxOpenConns    int
		ConnMaxLifetime int
	}

	redisConfig struct {
		Host     string
		Port     int
		Password string
	}
	jwtConfig struct {
		Secret string
		Expire int
	}
)

func init() {
	viper.AutomaticEnv()
	env := viper.Get("GO_ENV")
	if env == "production" {
		viper.SetConfigName("conf")
	} else {
		viper.SetConfigName("conf_dev")
	}

	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&item)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}

// Item return the loaded config content
func Item() Configuration {
	return item
}
