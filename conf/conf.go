package conf

import "github.com/jasontconnell/conf"

const SampleJson string = `{
    "proxyHost": "dev.example.dev",
    "proxyScheme": "https",
    "localHost": "example.local.dev",
    "localScheme": "https",
    "localKeyFile": "local.example.dev.key",
    "localCrtFile": "local.example.dev.crt",
    "localPort": 5454,
    "requestHeaders": {"Authorization": "Basic xxxxx"},
    "intercepts": [
        { "extension": "css", "mimeType": "text/css" },
        { "extension": "js", "mimeType": "application/javascript" }
    ],
    "localStartPath": "C:\\path\\to\\intercepted\\resources"
}`

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
}

type InterceptType struct {
	Extension string `json:"extension"`
	MimeType  string `json:"mimeType"`
}

func LoadConfig(filename string) Config {
	cfg := Config{
		ProxyScheme: "http",
		LocalScheme: "http",
	}

	conf.LoadConfig(filename, &cfg)
	return cfg
}
