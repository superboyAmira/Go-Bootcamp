package cfg

import (
	"encoding/xml"
	"os"
)

type Config struct {
	Port string `xml:"port"`
}

func LoadCfg(path string) (*Config, error) {
	cfg := &Config{}
	xmlData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(xmlData, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
} 