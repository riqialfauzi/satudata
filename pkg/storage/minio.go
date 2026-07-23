package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/satudata/backend/internal/config"
)

// StorageClient adalah wrapper untuk MinIO client.
type StorageClient struct {
	client  *minio.Client
	buckets BucketNames
}

// BucketNames menyimpan nama-nama bucket yang digunakan.
type BucketNames struct {
	Datasets  string
	Articles  string
	Documents string
}

// Init menginisialisasi koneksi ke MinIO.
func Init(cfg config.MinIOConfig) (*StorageClient, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
		Region: cfg.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	sc := &StorageClient{
		client: client,
		buckets: BucketNames{
			Datasets:  cfg.Buckets.Datasets,
			Articles:  cfg.Buckets.Articles,
			Documents: cfg.Buckets.Documents,
		},
	}

	// Buat buckets jika belum ada
	ctx := context.Background()
	buckets := []string{sc.buckets.Datasets, sc.buckets.Articles, sc.buckets.Documents}
	for _, bucket := range buckets {
		if err := sc.ensureBucket(ctx, bucket); err != nil {
			return nil, fmt.Errorf("failed to setup bucket %s: %w", bucket, err)
		}
	}

	log.Println("[Storage] Successfully connected to MinIO")
	return sc, nil
}

// ensureBucket membuat bucket jika belum ada.
func (s *StorageClient) ensureBucket(ctx context.Context, bucketName string) error {
	exists, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %w", err)
	}
	if !exists {
		err = s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
			Region: "us-east-1",
		})
		if err != nil {
			return fmt.Errorf("failed to create bucket %s: %w", bucketName, err)
		}
		log.Printf("[Storage] Created bucket: %s", bucketName)
	}
	return nil
}

// Upload mengupload file ke bucket yang ditentukan.
func (s *StorageClient) Upload(ctx context.Context, bucket, objectName string, reader io.Reader, fileSize int64, contentType string) (string, error) {
	_, err := s.client.PutObject(ctx, bucket, objectName, reader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Generate public URL
	url := fmt.Sprintf("%s/%s/%s", s.client.EndpointURL().String(), bucket, objectName)
	return url, nil
}

// Download mendownload file dari bucket.
func (s *StorageClient) Download(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
	reader, err := s.client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}
	return reader, nil
}

// Delete menghapus file dari bucket.
func (s *StorageClient) Delete(ctx context.Context, bucket, objectName string) error {
	return s.client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
}

// GeneratePresignedURL membuat URL sementara untuk upload/download.
func (s *StorageClient) GeneratePresignedURL(ctx context.Context, bucket, objectName string, expiry time.Duration) (string, error) {
	url, err := s.client.PresignedPutObject(ctx, bucket, objectName, expiry)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}
	return url.String(), nil
}

// GetBuckets mengembalikan daftar nama bucket.
func (s *StorageClient) GetBuckets() BucketNames {
	return s.buckets
}

// HealthCheck memeriksa koneksi MinIO.
func (s *StorageClient) HealthCheck(ctx context.Context) error {
	_, err := s.client.ListBuckets(ctx)
	return err
}
