package main

import (
    "bins/config"
    "bins/api"
)

func main() {
    cfg, _ := config.ReadEnv()
    apiClient := api.NewClient(cfg)  // Передали конфиг в API!
    // Теперь apiClient.CreateBin() будет использовать KEY
}