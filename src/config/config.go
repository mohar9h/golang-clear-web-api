package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Cors     CorsConfig
	Redis    RedisConfig
	Password PasswordConfig
}

type CorsConfig struct {
	AllowOrigins string
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

type RedisConfig struct {
	host               string
	port               string
	password           string
	db                 string
	minIdleConnections int
	poolSize           int
	poolTimeout        int
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}
	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to read config, %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func getConfigPath(env string) string {
	if env == "production" {
		return "config/config-production"
	} else {
		return "config/config-development"
	}
}
