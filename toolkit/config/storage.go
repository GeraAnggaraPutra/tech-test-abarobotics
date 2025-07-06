package config

import (
	"context"
	"os"
	"time"

	stg "abarobotics-test/src/handler/storage"
	"abarobotics-test/toolkit/logger"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

// Handle storage connection and open the connection to GCS.
func NewStorage() (s *stg.Storage, err error) {
	if bucketName := os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET_NAME"); bucketName != "" {
		var (
			client         *storage.Client
			credentialFile = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
			ctx            = context.Background()
		)

		if credentialFile != "" {
			client, err = storage.NewClient(ctx, option.WithCredentialsFile(credentialFile))
		} else {
			client, err = storage.NewClient(ctx)
		}

		if err != nil {
			logger.WithContext(ctx).Error(err, "error setup cloud storage")
			return
		}

		timeout, errT := time.ParseDuration(os.Getenv("GOOGLE_CLOUD_STORAGE_TIMEOUT"))
		if errT != nil {
			err = errors.Wrap(errT, "failed parse duration google cloud storage timeout")
			return
		}

		s = &stg.Storage{
			Client:     client,
			BucketName: bucketName,
			Timeout:    timeout,
		}
	}

	return
}
