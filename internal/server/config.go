package server

import (
	"os"
	"strconv"
)

type Config struct {
	Port           int
	WinstonBaseUrl string
	WinstonApiKey  string
}

func GetConfig() Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	conf := Config{
		Port:           port,
		WinstonBaseUrl: os.Getenv("WINSTON_API_URL"),
		WinstonApiKey:  os.Getenv("WINSTON_API_KEY"),
	}

	return conf
}
