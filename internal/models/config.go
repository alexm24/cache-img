package models

type HTTPServerConfig struct {
	BasePath string `yaml:"base_path"`
	Port     string `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type Config struct {
	HTTPServerConfig `yaml:"http_server_config"`
	DBConfig         `yaml:"db_config"`
}
