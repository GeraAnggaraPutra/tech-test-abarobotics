package storage

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"abarobotics-test/src/util"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

type Storage struct {
	Client     *storage.Client
	BucketName string
	Timeout    time.Duration
}

// Upload an object file to storage bucket.
func (stg Storage) UploadFile(ctx context.Context, file *multipart.FileHeader, filePath string) (err error) {
	src, err := file.Open()
	if err != nil {
		err = errors.Wrap(err, "file can't be opened")
		return
	}
	defer util.CloseBuffer(src)

	wc := stg.Client.Bucket(stg.BucketName).Object(filePath).NewWriter(ctx)
	defer util.CloseBuffer(wc)

	if _, err = io.Copy(wc, src); err != nil {
		err = errors.Wrap(err, "failed copy object file to google cloud storage")
		return
	}

	defer util.DiscardBuffer(src)

	return
}

// Get signed url.
func (stg Storage) GetSignedURL(filePath string) (url string, err error) {
	opts := &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(stg.Timeout),
	}

	url, err = stg.Client.Bucket(stg.BucketName).SignedURL(filePath, opts)
	if err != nil {
		err = errors.Wrap(err, "failed get signed url from google cloud storage")
		return
	}

	return
}

// Get blob stream an object file from storage bucket.
func (stg Storage) GetFile(ctx context.Context, filePath string) (rc *storage.Reader, err error) {
	rc, err = stg.Client.Bucket(stg.BucketName).Object(filePath).NewReader(ctx)
	if err != nil {
		err = errors.Wrap(err, "failed read object file from google cloud storage")
		return
	}

	defer util.CloseBuffer(rc)

	return
}

// Delete an object file from storage bucket.
func (stg Storage) DeleteFile(ctx context.Context, filePath string) (err error) {
	if err = stg.Client.Bucket(stg.BucketName).Object(filePath).Delete(ctx); err != nil {
		err = errors.Wrap(err, "failed delete object file from google cloud storage")
		return
	}

	return
}
