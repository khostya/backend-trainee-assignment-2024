package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		PG    PG    `yaml:"postgres"`
		CACHE CACHE `yaml:"cache"`
		HTTP  HTTP  `yaml:"http" env-prefix:"http"`
	}

	HTTP struct {
		Port           string        `env-required:"true" yaml:"port"`
		ReadTimeout    time.Duration `env-required:"true" yaml:"read_timeout"`
		WriteTimeout   time.Duration `env-required:"true" yaml:"write_timeout"`
		IdleTimeout    time.Duration `env-required:"true" yaml:"idle_timeout"`
		MaxHeaderBytes int           `env-required:"true" yaml:"max_header_bytes"`
	}

	CACHE struct {
		TTL         time.Duration `env-required:"true" yaml:"ttl"`
		MaxCost     int64         `env-required:"true" yaml:"max_cost"`
		NumCounters int64         `env-required:"true" yaml:"num_counters"`
	}

	PG struct {
		URL             string        `env:"PG_URL" env-default:"postgres://user:password@postgres:5432/postgres?sslmode=disable"`
		MaxOpenConns    int           `env-required:"true" yaml:"max_open_conns"`
		MaxIdleConns    int           `env-required:"true" yaml:"max_idle_conns"`
		ConnMaxIdleTime time.Duration `env-required:"true" yaml:"conn_max_idle_time"`
		ConnMaxLifetime time.Duration `env-required:"true" yaml:"conn_max_lifetime"`
	}
)

func NewConfig() (*Config, error) {
	configPath := "./config/config.yml"
	path, exists := os.LookupEnv("CONFIG_PATH")
	if exists {
		configPath = path
	}

	return NewConfigWithPath(configPath)
}

func NewConfigWithPath(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

func MustConfig() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
