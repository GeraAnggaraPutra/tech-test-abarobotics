package constant

import "time"

// middleware.
const (
	DefaultMdwHeaderToken         = "Authorization"
	DefaultMdwHeaderBearer        = "Bearer"
	DefaultMdwRateLimiter         = 20
	DefaultMdwRateLimiterDuration = time.Minute
	DefaultMdwTimeout             = 10 * time.Second
)

// pagination.
const (
	DefaultOrder = "created_at DESC"
	DefaultPage  = 1
	DefaultLimit = 10
)

const DefaultCacheExpireDuration = 24 * time.Hour

// transaction.
const (
	DefaultAdminFee             = 0.005
	DefaultInterestRatePerMonth = 0.03
)
