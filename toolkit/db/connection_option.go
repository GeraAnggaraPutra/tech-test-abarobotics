package db

import (
	"time"

	"abarobotics-test/toolkit/constant"

	"gorm.io/gorm/logger"
)

// ConnectionOption is db connection option.
type connectionOption struct {
	maxIdle           int
	maxOpen           int
	maxLifetime       time.Duration
	keepAliveInterval time.Duration
	gormOption
}

type gormOption struct {
	isLog         bool
	slowThreshold time.Duration
	level         logger.LogLevel
	ignoreErr     bool
	colorful      bool
}

// DefaultConnectionOption returns sensible conn setting.
func defaultConnectionOption() *connectionOption {
	return &connectionOption{
		maxIdle:           constant.DefaultDBMaxIdle,
		maxOpen:           constant.DefaultDBMaxOpen,
		maxLifetime:       constant.DefaultDBMaxLifetime,
		keepAliveInterval: constant.DefaultDBKeepAliveInterval,
		gormOption: gormOption{
			isLog:         constant.DefaultGormLog,
			slowThreshold: constant.DefaultGormLogSlowThreshold,
			level:         constant.DefaultGormLogLevel,
			ignoreErr:     constant.DefaultGormLogIgnoreErr,
			colorful:      constant.DefaultGormLogColorful,
		},
	}
}
