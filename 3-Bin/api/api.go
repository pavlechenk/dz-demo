package api

import "fmt"

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

func (c *Client) FetchData() (string, error) {
	fmt.Printf("[MOCK API] Запрос отправлен на %s\n", c.BaseURL)
	return `{"status": "ok", "data": "mock response"}`, nil
}
