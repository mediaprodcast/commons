package storage

import (
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/mediaprodcast/commons/env"
	storagePb "github.com/mediaprodcast/proto/genproto/go/storage/v1"
	"go.uber.org/zap"
)

type option func(*StorageService)

type StorageService struct {
	storageConfig *storagePb.StorageConfig
	endpoint      string
	directory     string
	retryCount    int
	retryDelay    time.Duration
	logger        *zap.Logger
	workers       int // Number of workers for parallel uploads
	client        *http.Client
}

func NewStorageService(options ...option) *StorageService {
	s := &StorageService{
		storageConfig: &storagePb.StorageConfig{
			Driver: storagePb.StorageDriver_FS,
			Fs: &storagePb.FileSystemConfig{
				DataPath: "./tmp",
			},
		},
		directory:  "",
		endpoint:   env.GetString("STORAGE_SERVICE_ADDR", "http://localhost:9500"),
		retryCount: 10,              // Default retry count
		retryDelay: 1 * time.Second, // Default retry delay
		logger:     zap.NewNop(),    // Default logger
		workers:    5,               // Default number of workers
	}

	for _, fn := range options {
		fn(s)
	}

	// Set client and access token
	s.client = s.GetHttpClient()

	return s
}

func (s *StorageService) GetHttpClient() *http.Client {
	return newClient(s.GetAccessToken, s.retryCount, s.retryDelay, s.logger)
}

// GetAccessToken returns an access token for the storage service.
func (s *StorageService) GetAccessToken() (string, error) {
	return EncodeStorageConfig(s.storageConfig)
}

// GetAbsolutePath returns an absolute representation of the path.
func (s *StorageService) GetAbsolutePath(name string) string {
	if isValidURL(s.endpoint) {
		baseURL, err := url.Parse(s.endpoint)
		if err == nil {
			return baseURL.String() + "/" + path.Join(s.directory, name)
		}
	}

	return path.Join(s.endpoint, s.directory, name)
}

// isValidURL checks if a string is a valid HTTP or HTTPS URL.
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https")
}
