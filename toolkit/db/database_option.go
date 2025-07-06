package db

import (
	"os"
	"strconv"

	"abarobotics-test/toolkit/util"

	"github.com/pkg/errors"
)

type databaseOption struct {
	driver       string
	host         string
	port         int
	username     string
	password     string
	schema       string
	sslmode      string
	usePrivate   string
	instanceName string
	*connectionOption
}

func newDatabaseOption() (*databaseOption, error) {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	schema := os.Getenv("DB_SCHEMA")
	sslmode := os.Getenv("DB_SSLMODE")
	usePrivate := os.Getenv("PRIVATE_IP")
	instanceName := os.Getenv("INSTANCE_CONNECTION_NAME")

	if portStr == "" {
		portStr = "0"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse int on port db env : %s", portStr)
	}

	if host == "" {
		return nil, errors.Wrapf(errors.New("invalid data source host or port"), "db: host=%s port=%d", host, port)
	}

	conn := defaultConnectionOption()
	conn.maxIdle = util.ParseInt(conn.maxIdle, os.Getenv("DB_MAX_IDLE_CONN"))
	conn.maxOpen = util.ParseInt(conn.maxOpen, os.Getenv("DB_MAX_OPEN_CONN"))
	conn.maxLifetime = util.ParseDuration(conn.maxLifetime, os.Getenv("DB_MAX_LIFETIME_CONN"))
	conn.keepAliveInterval = util.ParseDuration(conn.keepAliveInterval, os.Getenv("DB_KEEP_ALIVE_INTERVAL_CONN"))

	return &databaseOption{
		driver:           driver,
		host:             host,
		port:             port,
		username:         username,
		password:         password,
		schema:           schema,
		sslmode:          sslmode,
		connectionOption: conn,
		usePrivate:       usePrivate,
		instanceName:     instanceName,
	}, nil
}
