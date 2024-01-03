package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeq = "{password}"
	password     = "note-service-password"
)

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

type Config struct {
	HTTP HTTP `json:"http"`
	GRPC GRPC `json:"grpc"`
	DB   DB   `json:"db"`
}

func New(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err = json.Unmarshal(file, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (g *GRPC) GetAddress() string {
	return net.JoinHostPort(g.Host, g.Port)
}

func (h *HTTP) GetAddress() string {
	return net.JoinHostPort(h.Host, h.Port)
}

func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
	poolConfig, err := pgxpool.ParseConfig(strings.ReplaceAll(c.DB.DSN, dbPassEscSeq, password))
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}
