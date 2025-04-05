package caddy

import (
	"encoding/json"
	"os"
)

type CaddyConfig struct {
	Apps struct {
		HTTP struct {
			Servers map[string]Server `json:"servers"`
		} `json:"http"`
	} `json:"apps"`
}

type Server struct {
	Listen []string `json:"listen"`
	Routes []Route  `json:"routes"`
	Name   string   `json:"-"`
}

type Route struct {
	Handle []Handler `json:"handle"`
}

type Handler struct {
	Handler   string     `json:"handler"`
	Upstreams []Upstream `json:"upstreams,omitempty"`
}

type Upstream struct {
	Dial string `json:"dial"`
}

func LoadConfig() (*CaddyConfig, error) {
	raw, err := os.ReadFile("/etc/caddy/caddy.json")
	if err != nil {
		return nil, err
	}
	var config CaddyConfig
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func SaveConfig(cfg *CaddyConfig) error {
	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("/etc/caddy/caddy.json", out, 0644)
}
