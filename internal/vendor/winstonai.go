package vendor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Service interface {
	DetectAiFromText(r *http.Request) (map[string]interface{}, error)
	TestGetEndpoint(r *http.Request) (map[string]interface{}, error)
	TestPostEndpoint(r *http.Request) (map[string]interface{}, error)
}

type service struct {
	baseUrl string
	apiKey  string
}

var (
	baseUrl = os.Getenv("WINSTON_API_URL")
	apiKey  = os.Getenv("WINSTON_API_KEY")
)

func New() Service {
	winstonai := &service{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
	return winstonai
}

func (s service) DetectAiFromText(r *http.Request) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s%s", s.baseUrl, "/v2/ai-content-detection")
	bearer := fmt.Sprintf("Bearer %s", s.apiKey)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get content for detection: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url,
		bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return data, nil
}

func (s service) TestGetEndpoint(r *http.Request) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s", "https://jsonplaceholder.typicode.com/todos/1")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return data, nil
}

func (s service) TestPostEndpoint(r *http.Request) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s", "https://jsonplaceholder.typicode.com/posts")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get content for detection: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url,
		bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return data, nil
}
