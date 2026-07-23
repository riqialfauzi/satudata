package service

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/satudata/backend/pkg/storage"
)

// StorageService adalah implementasi dari StorageServiceInterface.
type StorageService struct {
	client *storage.StorageClient
}

// NewStorageService membuat instance baru StorageService.
func NewStorageService(client *storage.StorageClient) *StorageService {
	return &StorageService{
		client: client,
	}
}

// UploadDataset mengupload file dataset.
func (s *StorageService) UploadDataset(ctx context.Context, file UploadFileRequest) (string, error) {
	if s.client == nil {
		return "", fmt.Errorf("storage not available")
	}

	objectName := fmt.Sprintf("datasets/%s/%s", uuid.New().String(), file.FileName)
	reader := bytes.NewReader(file.Data)

	url, err := s.client.Upload(ctx, s.client.GetBuckets().Datasets, objectName, reader, file.FileSize, file.ContentType)
	if err != nil {
		return "", fmt.Errorf("failed to upload dataset: %w", err)
	}

	return url, nil
}

// UploadArticleImage mengupload image artikel.
func (s *StorageService) UploadArticleImage(ctx context.Context, file UploadFileRequest) (string, error) {
	if s.client == nil {
		return "", fmt.Errorf("storage not available")
	}

	objectName := fmt.Sprintf("articles/%s/%s", uuid.New().String(), file.FileName)
	reader := bytes.NewReader(file.Data)

	url, err := s.client.Upload(ctx, s.client.GetBuckets().Articles, objectName, reader, file.FileSize, file.ContentType)
	if err != nil {
		return "", fmt.Errorf("failed to upload article image: %w", err)
	}

	return url, nil
}

// UploadStandardDoc mengupload dokumen standard.
func (s *StorageService) UploadStandardDoc(ctx context.Context, file UploadFileRequest) (string, error) {
	if s.client == nil {
		return "", fmt.Errorf("storage not available")
	}

	objectName := fmt.Sprintf("documents/%s/%s", uuid.New().String(), file.FileName)
	reader := bytes.NewReader(file.Data)

	url, err := s.client.Upload(ctx, s.client.GetBuckets().Documents, objectName, reader, file.FileSize, file.ContentType)
	if err != nil {
		return "", fmt.Errorf("failed to upload standard doc: %w", err)
	}

	return url, nil
}

// DeleteFile menghapus file dari storage.
func (s *StorageService) DeleteFile(ctx context.Context, url string) error {
	if s.client == nil {
		return fmt.Errorf("storage not available")
	}

	// Parse URL to get bucket and object name
	// Format: http://endpoint/bucket/object-name
	bucket, objectName := parseStorageURL(url, s.client.GetBuckets())
	if bucket == "" || objectName == "" {
		return fmt.Errorf("invalid file URL")
	}

	return s.client.Delete(ctx, bucket, objectName)
}

// GeneratePresignedURL membuat presigned URL untuk upload sementara.
func (s *StorageService) GeneratePresignedURL(ctx context.Context, key string, expiry int32) (string, error) {
	if s.client == nil {
		return "", fmt.Errorf("storage not available")
	}

	duration := time.Duration(expiry) * time.Minute
	return s.client.GeneratePresignedURL(ctx, s.client.GetBuckets().Datasets, key, duration)
}

// parseStorageURL mengekstrak bucket dan object name dari URL.
func parseStorageURL(url string, buckets storage.BucketNames) (string, string) {
	// Simple URL parsing - extract bucket and key from the URL path
	bucketNames := []string{buckets.Datasets, buckets.Articles, buckets.Documents}

	for _, bucket := range bucketNames {
		prefix := fmt.Sprintf("/%s/", bucket)
		// Look for the bucket name in the URL
		for i := 0; i <= len(url)-len(prefix); i++ {
			if url[i:i+len(prefix)] == prefix {
				objectName := url[i+len(prefix):]
				return bucket, objectName
			}
		}
	}

	return "", ""
}
