package driver

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Name     string `env:"DB_SCHEMA"`
	Adapter  string `env:"DB_DRIVER"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	UserDB   string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type ServerConfig struct {
	ServiceName string `env:"SERVICE_NAME"`
	ServicePort string `env:"PORT"`
	ServiceHost string `env:"SERVICE_HOST"`
	DB          Database
}

var Config ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Warn().Msg("Cannot find .env file. OS Environments will be used")
	}
	err = env.Parse(&Config)
	err = env.Parse(&Config.DB)

	return err
}

//func IsDevelopmentMode() bool {
//	return os.Getenv("ENV") == "DEVELOPMENT"
//}
//
//func IsDebugMode() bool {
//	return os.Getenv("ENV") == "DEBUG" || IsDevelopmentMode()
//}

