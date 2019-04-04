/**
 * @author Jose Nidhin
 */
package system

import "sync"

type Config struct {
	Log *LogConfig `json:"log"`

	Server *ServerConfig `json:"server"`
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
		}
	})

	return instance
}
