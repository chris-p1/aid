package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/v1/ai/evaltext", s.aiTextDetectionHandler)
	mux.HandleFunc("/v1/test/get", s.testGetEndpointHandler)
	mux.HandleFunc("/v1/test/post", s.testPostEndpointHandler)

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) aiTextDetectionHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s%s", s.winston.BaseUrl, "/v2/ai-content-detection")
	body, err := io.ReadAll(r.Body)

	resp, err := s.winston.Post(url, body, s.winston.ApiKey)
	if err != nil {
		log.Printf("Failed to get response: %v", err)
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) testGetEndpointHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/todos/1", s.tester.BaseUrl)

	resp, err := s.tester.Get(url)
	if err != nil {
		log.Printf("Failed to get response: %v", err)
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) testPostEndpointHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", s.tester.BaseUrl)
	body, err := io.ReadAll(r.Body)

	resp, err := s.winston.Post(url, body, "")
	if err != nil {
		log.Printf("Failed to get response: %v", err)
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
