package db

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	myx "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysql(opt *databaseOption) (dbx *sqlx.DB, db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", opt.username, opt.password, opt.host, opt.port, opt.schema)

	if opt.instanceName != "" {
		dsn = fmt.Sprintf("%s:%s@cloudsqlconn(%s:%d)/%s?parseTime=true", opt.username, opt.password, opt.host, opt.port, opt.schema)

		var (
			opts []cloudsqlconn.DialOption
			d    *cloudsqlconn.Dialer
		)

		if opt.usePrivate != "" {
			opts = append(opts, cloudsqlconn.WithPrivateIP())
		}

		d, err = cloudsqlconn.NewDialer(context.Background())
		if err != nil {
			err = errors.Wrap(err, "cloudsql: failed to make dialer")
			return
		}

		myx.RegisterDialContext("cloudsqlconn", func(ctx context.Context, _ string) (net.Conn, error) {
			return d.Dial(ctx, opt.instanceName, opts...)
		})
	}

	dbx, err = openSQL("mysql", dsn, opt)
	if err != nil {
		err = errors.Wrap(err, "mysql: failed to open connection")
		return
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbx.DB,
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             opt.slowThreshold,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		err = errors.Wrap(err, "gorm: failed to open connection")
		return
	}

	log.Printf("successfully connected to mysql %s:%d", opt.host, opt.port)

	go keepAlive(dbx, opt.driver, opt.schema, opt.keepAliveInterval)

	return
}
