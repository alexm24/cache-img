package models

type HTTPServerConfig struct {
	BasePath string `yaml:"base_path"`
	Port     string `yaml:"port"`
}

type Config struct {
	HTTPServer HTTPServerConfig `yaml:"http_server"`
}
