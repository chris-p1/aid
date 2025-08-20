package winstonapi

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// // Service represents a service that interacts with winstonAI
// type Service interface {
// 	// returns a map of AI detection results
// 	AITextDetection() map[string]interface{}

// }

// type service struct {
// 	url *string
// 		apiKey *string		
// }

// var (
// 		url = os.Getenv("WINSTON_API_URL")
// 		apiKey = os.Getenv("WINSTON_API_KEY")
// 	winstance *service
// )

// func New() Service {

// 	winstance = &service{
// 		url: url,
// 			apiKey: apiKey,
// 	}
// 	return winstance
// }

// func (s *service) AITextDetection() map[string]interface{} {
		
// }
