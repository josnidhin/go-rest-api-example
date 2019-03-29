/**
 * @author Jose Nidhin
 */
package config

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

func Load() *Config {
	return &Config{
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
}
