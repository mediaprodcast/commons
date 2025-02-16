package storage

import (
	"time"

	storagePb "github.com/mediaprodcast/proto/genproto/go/storage/v1"
	"go.uber.org/zap"
)

// WithEndpoint sets the endpoint for the StorageService.
// This is useful for configuring the target storage service address.
func WithEndpoint(endpoint string) option {
	return func(s *StorageService) {
		s.endpoint = endpoint
	}
}

// WithDirectory sets the directory for the StorageService.
// This is useful for specifying a subdirectory within the storage service.
func WithDirectory(directory string) option {
	return func(s *StorageService) {
		s.directory = directory
	}
}

// WithStorageConfig sets the storage configuration for the StorageService.
// This is useful for customizing the storage driver and its settings.
func WithStorageConfig(storageConfig *storagePb.StorageConfig) option {
	return func(s *StorageService) {
		s.storageConfig = storageConfig
	}
}

// WithRetryCount sets the maximum number of retries for upload operations.
// This is useful for handling transient errors during uploads.
func WithRetryCount(count int) option {
	return func(s *StorageService) {
		s.retryCount = count
	}
}

// WithRetryDelay sets the delay between retries for upload operations.
// This is useful for implementing exponential backoff or fixed delays between retries.
func WithRetryDelay(delay time.Duration) option {
	return func(s *StorageService) {
		s.retryDelay = delay
	}
}

// WithLogger sets a custom logger for the StorageService.
// This is useful for integrating with existing logging systems or customizing log output.
func WithLogger(logger *zap.Logger) option {
	return func(s *StorageService) {
		s.logger = logger
	}
}

// WithWorkers sets the number of workers for parallel uploads.
func WithWorkers(count int) option {
	return func(s *StorageService) {
		s.workers = count
	}
}
