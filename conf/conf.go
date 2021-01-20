package conf

import "github.com/jasontconnell/conf"

type Headers map[string]string

type Config struct {
	ProxyScheme    string          `json:"proxyScheme"`
	ProxyHost      string          `json:"proxyHost"`
	LocalPort      int             `json:"localPort"`
	RequestHeaders Headers         `json:"requestHeaders"`
	Intercepts     []InterceptType `json:"intercepts"`
	LocalStartPath string          `json:"localStartPath"`
}

type InterceptType struct {
	Extension string `json:"extension"`
	MimeType  string `json:"mimeType"`
}

func LoadConfig(filename string) Config {
	cfg := Config{}

	conf.LoadConfig(filename, &cfg)
	return cfg
}
