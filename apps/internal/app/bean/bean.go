package bean

import (
	"gorm.io/gorm"
	"your_spotify/pkg/out"
)

type Bean struct {
	logger out.Logger
	config Config
	db     *gorm.DB
}

func (b Bean) Logger() out.Logger {
	return b.logger
}

func (b Bean) Config() Config {
	return b.config
}

func (b Bean) Db() *gorm.DB {
	return b.db
}

var bean *Bean

// InitBean 순서 보장을 위해 파람을 받는다.
func InitBean(config *Config) *Bean {
	if bean != nil {
		return bean
	}
	db := initDB(config)

	bean = &Bean{
		logger: out.NewLogger(config.LogLevel()),
		config: *config,
		db:     db,
	}
	return bean
}

func GetBean() *Bean {
	return bean
}
