package config

type GlobalConfig struct {
	HTTP             HTTPConfig       `yaml:"http"`
	BookService      BookService      `yaml:"bookservice"`
	HttpClientConfig HttpClientConfig `yaml:"httpclientconfig"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type BookService struct {
	Address string `yaml:"address"`
}

type HttpClientConfig struct {
	TimeoutMS           int `yaml:"timeoutms"`
	MaxIdleConns        int `yaml:"maxidleconns"`
	MaxIdleConnsPerHost int `yaml:"maxidleconnsperhost"`
	MaxConnsPerHost     int `yaml:"maxconnsperhost"`
	IdleConnTimeoutSec  int `yaml:"idleconntimeoutsec"`
}
