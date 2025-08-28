package api

type Api interface {
	Get(url string) (map[string]any, error)
	Post(url string, body []byte, auth string) (map[string]any, error)
}
