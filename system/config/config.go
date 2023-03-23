/**
 * @author Jose Nidhin
 */
package config

import (
	"os"

	"github.com/josnidhin/go-rest-api-example/system/logger"
	"github.com/josnidhin/go-rest-api-example/system/pgdb"
)

type Config struct {
	Log *logger.LogConfig `json:"log"`

	Server *ServerConfig `json:"server"`

	PG *pgdb.PGConfig `json:"pg"`
}

type ServerConfig struct {
	HTTP *HTTPServerConfig `json:"http"`
}

type HTTPServerConfig struct {
	Port int `json:"port"`
}

func New() *Config {
	config := &Config{
		Log: &logger.LogConfig{
			Level:       "debug",
			Development: true,
		},
		Server: &ServerConfig{
			HTTP: &HTTPServerConfig{
				Port: 8080,
			},
		},
		PG: &pgdb.PGConfig{
			Host:     os.Getenv("PG_DEFAULT_HOST"),
			User:     os.Getenv("PG_DEFAULT_USER"),
			Password: os.Getenv("PG_DEFAULT_PASSWORD"),
			Port:     os.Getenv("PG_DEFAULT_PORT"),
			DBName:   os.Getenv("PG_DEFAULT_DBNAME"),
			SSLMode:  "disable",
		},
	}

	return config
}
