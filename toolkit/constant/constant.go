package constant

import (
	"time"

	"gorm.io/gorm/logger"
)

// echo runtime.
const (
	DefaultAppPort            = 8000
	DefaultAppPrometheus      = false
	DefaultAppShutdownTimeout = 200 * time.Millisecond
	DefaultAppShutdownWait    = 200 * time.Millisecond
)

// logger option.
const (
	DefaultLogLevel              = 0
	DefaultLogCallerSkipFrame    = 3
	DefaultLogCallerMaxDirectory = 5
	DefaultSentryLevel           = 3
)

// db connection.
const (
	DefaultDBMaxIdle           = 20
	DefaultDBMaxOpen           = 100
	DefaultDBMaxLifetime       = 10 * time.Second
	DefaultDBKeepAliveInterval = 3 * time.Minute

	// gorm option.
	DefaultGormLog              = true
	DefaultGormLogSlowThreshold = 200 * time.Millisecond
	DefaultGormLogLevel         = logger.Warn
	DefaultGormLogIgnoreErr     = false
	DefaultGormLogColorful      = true
)

// cache connection.
const (
	DefaultCacheMaxRetries        = 3
	DefaultCacheMinRetryBackoff   = 8 * time.Millisecond
	DefaultCacheMaxRetryBackoff   = 512 * time.Millisecond
	DefaultCacheDialTimeout       = 5 * time.Second
	DefaultCacheReadTimeout       = 3 * time.Second
	DefaultCacheWriteTimeout      = 3 * time.Second
	DefaultCachePoolFIFO          = true
	DefaultCachePoolSize          = 10
	DefaultCachePoolTimeout       = 4 * time.Second
	DefaultCacheMaxIdleConns      = 0
	DefaultCacheMaxActiveConns    = 0
	DefaultCacheConnMaxIdleTime   = 30 * time.Minute
	DefaultCacheConnMaxLifeTime   = 0
	DefaultCacheIdentitySuffix    = ""
	DefaultCacheKeepAliveInterval = 3 * time.Minute
)
