/**
 * @author Jose Nidhin
 */
package config

type Config struct {
	Log *LogConfig `json: log`

	Server *ServerConfig `json: server`
}

type LogConfig struct {
	Level string `json: level`
}

type ServerConfig struct {
	HTTP *HTTPServerConfig `json: http`
}

type HTTPServerConfig struct {
	Port int `json: port`
}

func Load() *Config {
	return &Config{
		Log: &LogConfig{
			Level: "debug",
		},
		Server: &ServerConfig{
			HTTP: &HTTPServerConfig{
				Port: 8080,
			},
		},
	}
}
