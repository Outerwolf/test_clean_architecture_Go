package config

type Configuration struct {
	Environment string     `yaml:"environment"`
	AppName     string     `yaml:"appName"`
	HttpServer  HttpServer `yaml:"HttpServer"`
}

type HttpServer struct {
	Port string `yaml:"Port"`
}
