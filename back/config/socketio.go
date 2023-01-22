package config

import (
	"net/url"
)

type SocketIO struct {
	URL       string `mapstructure:"url"`
	ParsedURL *url.URL
}

func (config *SocketIO) Load() error {
	parsed, err := url.Parse(config.URL)
	if err != nil {
		return err
	}
	config.ParsedURL = parsed
	return nil
}
