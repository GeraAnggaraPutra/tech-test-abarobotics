package config

import (
	"os"
	"time"

	"abarobotics-test/toolkit/constant"
	"abarobotics-test/toolkit/util"

	"github.com/iancoleman/strcase"
)

// RuntimeConfig defines echo REST API runtime config.
type RuntimeConfig struct {
	Name                    string        `json:"name"`
	Host                    string        `json:"host"`
	Port                    int           `json:"port"`
	Prometheus              bool          `json:"prometheus"`
	ShutdownTimeoutDuration time.Duration `json:"shutdown_timeout_duration"`
	ShutdownWaitDuration    time.Duration `json:"shutdown_wait_duration"`
}

func NewRuntimeConfig() *RuntimeConfig {
	r := RuntimeConfig{}

	r.Name = os.Getenv("APP_NAME")
	r.Host = os.Getenv("APP_HOST")
	r.Port = util.ParseInt(constant.DefaultAppPort, os.Getenv("APP_PORT"))
	r.Prometheus = util.ParseBool(constant.DefaultAppPrometheus, os.Getenv("APP_PROMETHEUS"))
	r.ShutdownTimeoutDuration = constant.DefaultAppShutdownTimeout
	r.ShutdownWaitDuration = constant.DefaultAppShutdownWait
	r.Name = strcase.ToSnake(r.Name)

	return &r
}
