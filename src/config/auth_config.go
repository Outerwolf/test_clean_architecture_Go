package config

type Configuration struct {
	AppName    string     `yaml:"appName"`
	HttpServer HttpServer `yaml:"HttpServer"`
}

type HttpServer struct {
	Port string `yaml:"Port"`
}
