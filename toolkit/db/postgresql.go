package db

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDatabase - create & validate postgres connection given certain db.Option
// the caller have the responsibility to close the *sqlx.DB when succeed.
func NewPostgresql(opt *databaseOption) (dbx *sqlx.DB, db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", opt.host, opt.port, opt.username, opt.password, opt.schema, opt.sslmode)

	if opt.instanceName != "" {
		dsn = fmt.Sprintf("user=%s password=%s database=%s", opt.username, opt.password, opt.schema)
	}

	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		err = errors.Wrap(err, "postgres: failed parse connection config")
		return
	}

	if opt.instanceName != "" {
		var (
			opts []cloudsqlconn.Option
			d    *cloudsqlconn.Dialer
		)

		if opt.usePrivate != "" {
			opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
		}

		d, err = cloudsqlconn.NewDialer(context.Background(), opts...)
		if err != nil {
			err = errors.Wrap(err, "cloudsql: failed to make dialer")
			return
		}

		config.DialFunc = func(ctx context.Context, _, _ string) (net.Conn, error) {
			return d.Dial(ctx, opt.instanceName)
		}
	}

	dbx, err = openSQL("pgx", stdlib.RegisterConnConfig(config), opt)
	if err != nil {
		err = errors.Wrap(err, "postgres: failed to open connection")
		return
	}

	var cfgLogger logger.Interface

	if opt.isLog {
		cfgLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             opt.slowThreshold,
				LogLevel:                  opt.level,
				IgnoreRecordNotFoundError: opt.ignoreErr,
				Colorful:                  opt.colorful,
			},
		)
	}

	db, err = gorm.Open(postgres.New(postgres.Config{Conn: dbx.DB}), &gorm.Config{Logger: cfgLogger})
	if err != nil {
		err = errors.Wrap(err, "gorm: failed to open connection")
		return
	}

	log.Printf("successfully connected to postgresql %s:%d", opt.host, opt.port)

	go keepAlive(dbx, opt.driver, opt.schema, opt.keepAliveInterval)

	return
}
