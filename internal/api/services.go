package api

import (
	_ "github.com/joho/godotenv/autoload"
)

type Service struct {
	url       string
	endpoints map[string]string
	key       string
}

func NewService() *Service {

	newService := &Service{
		url:       url,
		endpoints: endpoints,
		key:       key,
	}

	return newService
}
