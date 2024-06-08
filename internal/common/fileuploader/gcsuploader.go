package fileuploader

import (
	"cloud.google.com/go/storage"
	"context"
)

type IStorage interface {
	UploadFile(file []byte, path string) (string, error)
}

// GCSUploader is an adapter that implements the Uploader interface using Google Cloud Storage
type GCSUploader struct {
	BucketName string
	Client     *storage.Client
}

// NewGCSUploader creates a new instance of GCSUploader
func NewGCSUploader(bucketName string) (*GCSUploader, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &GCSUploader{
		BucketName: bucketName,
		Client:     client,
	}, nil
}

// UploadFile uploads a file to Google Cloud Storage and returns the URL
func (u *GCSUploader) UploadFile(file []byte, path string) (string, error) {
	ctx := context.Background()
	wc := u.Client.Bucket(u.BucketName).Object(path).NewWriter(ctx)
	if _, err := wc.Write(file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	return "https://storage.googleapis.com/" + u.BucketName + "/" + path, nil
}
