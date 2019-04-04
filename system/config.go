/**
 * @author Jose Nidhin
 */
package system

import (
	"os"
	"sync"
)

type Config struct {
	Log *LogConfig `json:"log"`

	Server *ServerConfig `json:"server"`

	Database *DatabaseConfig `json:"database"`
}

type LogConfig struct {
	Level       string `json:"level"`
	Development bool   `json:development`
}

type ServerConfig struct {
	HTTP *HTTPServerConfig `json:"http"`
}

type HTTPServerConfig struct {
	Port int `json:"port"`
}

type DatabaseConfig struct {
	PG *PGConfigs `json:"pg"`
}

type PGConfigs struct {
	Default *PGConfig `json:"default"`
}

type PGConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	DBName string `json:"dbname"`
}

var instance *Config
var configOnce sync.Once

func ConfigInstance() *Config {
	configOnce.Do(func() {
		instance = &Config{
			Log: &LogConfig{
				Level:       "debug",
				Development: true,
			},
			Server: &ServerConfig{
				HTTP: &HTTPServerConfig{
					Port: 8080,
				},
			},
			Database: &DatabaseConfig{
				PG: &PGConfigs{
					Default: &PGConfig{
						Host:     os.Getenv("PG_DEFAULT_HOST"),
						User:     os.Getenv("PG_DEFAULT_USER"),
						Password: os.Getenv("PG_DEFAULT_PASSWORD"),
						Port:     os.Getenv("PG_DEFAULT_PORT"),
						DBName: os.Getenv("PG_DEFAULT_DBNAME"),
					},
				},
			},
		}
	})

	return instance
}
