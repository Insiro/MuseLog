package bean

import "os"

type Config struct {
	dbUrl    string
	baseURL  string
	logLevel string
}

func (c Config) DBUrl() string {
	return c.dbUrl
}
func (c Config) BaseURL() string {
	return c.baseURL
}
func (c Config) LogLevel() string {
	return c.logLevel
}

var config *Config

func InitConfig() *Config {
	if config != nil {
		return config
	}

	//TODO: load Config From Env
	config := &Config{}

	return config
}

// getEnv 함수는 환경 변수를 가져오고, 기본값을 반환합니다.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetConfig() *Config {
	return config
}
