package caddy

import (
	"encoding/json"
	"bytes"
	"fmt"
	"net/http"
)

func PushConfig(cfg *CaddyConfig) error {
	out, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:2019/load", "application/json", bytes.NewBuffer(out))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("caddy responded with status: %s", resp.Status)
	}
	return nil
}
