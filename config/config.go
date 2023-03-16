package config

import (
	"github.com/joho/godotenv"
	"os"
)

const envName = ".deploy/.env"

type Config struct{}

func NewConfig() *Config {
	if err := godotenv.Load(envName); err != nil {
		panic(err)
	}
	return &Config{}
}

func (c *Config) GetSelfPort() string {
	return os.Getenv("SELF_PORT")
}

func (c *Config) GetAreaPort() string {
	return os.Getenv("AREA_PORT")
}

func (c *Config) GetAuthPort() string {
	return os.Getenv("AUTH_PORT")
}
