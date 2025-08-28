package server

import (
	"aid/internal/service"
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port    int
	winston service.WinstonAPI
	tester  service.TestAPI
}

func NewServer() *http.Server {

	conf := GetConfig()
	NewServer := &Server{
		port: conf.Port,
		winston: service.WinstonAPI{
			BaseUrl: conf.WinstonBaseUrl,
			ApiKey:  conf.WinstonApiKey,
		},
		tester: service.TestAPI{
			BaseUrl: "https://jsonplaceholder.typicode.com",
		},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
