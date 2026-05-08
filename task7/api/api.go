package api

import "bins/config"
		

type Client struct {
    apiKey string
    baseURL string
}

func NewClient(cfg *config.Env) *Client {
    return &Client{
        apiKey: cfg.Key,
        baseURL: "https://api.jsonbin.io/v3",
    }
}

// func (c *Client) CreateBin(data []byte) error {
//     req, _ := http.NewRequest("POST", c.baseURL+"/b", bytes.NewBuffer(data))
//     req.Header.Set("X-Master-Key", c.apiKey) // Используем KEY здесь!
//     // ... остальная логика
//     return nil
// }