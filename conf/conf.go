package conf

import "github.com/jasontconnell/conf"

type Headers map[string]string

type Config struct {
	ProxyScheme    string          `json:"proxyScheme"`
	ProxyHost      string          `json:"proxyHost"`
	LocalHost      string          `json:"localHost"`
	LocalScheme    string          `json:"localScheme"`
	LocalKeyFile   string          `json:"localKeyFile"`
	LocalCrtFile   string          `json:"localCrtFile"`
	LocalPort      int             `json:"localPort"`
	RequestHeaders Headers         `json:"requestHeaders"`
	Intercepts     []InterceptType `json:"intercepts"`
	LocalStartPath string          `json:"localStartPath"`
	OverwriteHost  bool            `json:"overwriteHost"`
}

type InterceptType struct {
	Extension string `json:"extension"`
	MimeType  string `json:"mimeType"`
}

func LoadConfig(filename string) Config {
	cfg := Config{
		ProxyScheme:   "http",
		LocalScheme:   "http",
		OverwriteHost: true,
	}

	conf.LoadConfig(filename, &cfg)
	return cfg
}
