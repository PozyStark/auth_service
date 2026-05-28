package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	DefaultConfigFilePath = "./configs/auth_service/config.yaml"
)

type Config struct {
	Env                 string      `yaml:"env" env-default:"dev"`
	TablePrefix         string      `yaml:"table_prefix" env-default:"auth_service"`
	EnableAutomigration bool        `yaml:"enable_automigration" env-default:"false"`
	Loglevel            string      `yaml:"loglevel" env-default:"debug"`
	DbConfig            DbConfig    `yaml:"db_config"`
	HTTPServer          HTTPServer  `yaml:"http_server"`
	Crypt               Crypt       `yaml:"crypt"`
	Token               Token       `yaml:"token"`
}

type DbConfig struct {
	DbHost                string        `yaml:"db_host" env-required:"true"`
	DbPort                int           `yaml:"db_port" env-required:"true"`
	DbName                string        `yaml:"db_name" env:"DB_NAME" env-required:"true"`
	DbUser                string        `yaml:"db_user" env:"DB_USER" env-required:"true"`
	DbPassword            string        `yaml:"db_password" env:"DB_PASSWORD" env-required:"true"`
	SslMode               string        `yaml:"ssl_mode" env:"SSL_MODE" env-required:"true" env-default:"disabled"`
	MaxOpenConnections    int           `yaml:"max_open_connections" env-required:"true"`
	MaxIdleConnections    int           `yaml:"max_idle_connections" env-required:"true"`
	MaxConnectionLifetime time.Duration `yaml:"max_connection_lifetime" env-required:"true"`
	MaxConnectionIdletime time.Duration `yaml:"max_connection_idletime" env-required:"true"`
}

type HTTPServer struct {
	HttpAddress     string        `yaml:"http_address" env-default:"localhost"`
	HttpPort        int           `yaml:"http_port" env-default:"8080"`
	HttpTimeout     time.Duration `yaml:"http_timeout" env-default:"5s"`
	HttpIdleTimeout time.Duration `yaml:"http_idle_timeout" env-default:"60s"`
}

type Crypt struct {
	Times     uint32 `yaml:"times" env-default:"1" env-required:"true"`
	Memory    uint32 `yaml:"memory" env-default:"64" env-required:"true"`
	Threads   uint8  `yaml:"threads" env-default:"1" env-required:"true"`
	KeyLength uint32 `yaml:"key_length" env-default:"32" env-required:"true"`
	Salt      string `yaml:"salt" env:"SALT" env-required:"true"`
}

type Token struct {
	SecretRefreshToken   string        `yaml:"secret_refresh_token" env:"SECRET_REFRESH_TOKEN" env-required:"true"`
	SecretAccessToken    string        `yaml:"secret_access_token" env:"SECRET_ACCESS_TOKEN" env-required:"true"`
	AccessTokenLifetime  time.Duration `yaml:"access_token_lifetime" env-default:"3600s" env-requred:"true"`
	RefreshTokenLifetime time.Duration `yaml:"refresh_token_lifetime" env-default:"3600s" env-requred:"true"`
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Info .env file is not set")
	}

	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		log.Printf("Info enviroment CONFIG_FILE_PATH is not set")
		configFilePath = DefaultConfigFilePath
	}

	_, err = os.Stat(configFilePath)
	if err != nil {
		log.Fatalf("Error config file not exist: %v", err)
	}

	var cfg Config

	err = cleanenv.ReadConfig(configFilePath, &cfg)
	if err != nil {
		log.Fatalf("Error read config file: %v", err)
	}

	return &cfg
}
