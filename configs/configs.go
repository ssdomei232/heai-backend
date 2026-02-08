package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	MYSQL          MYSQL    `json:"mysql"`
	GsraiToken     string   `json:"grsai_token"`
	BaseURL        string   `json:"base_url"` // like: yourdomain.com
	TrustedOrigins []string `json:"trusted_origins"`
}

type MYSQL struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"database"`
	Port     int    `json:"port"`
}

func GetConfig() (config Configs, err error) {
	content, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
