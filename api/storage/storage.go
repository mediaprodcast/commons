package storage

import (
	"net/url"
	"path"

	"github.com/mediaprodcast/commons/env"
	storagePb "github.com/mediaprodcast/proto/genproto/go/storage/v1"
)

type StorageService struct {
	storageConfig *storagePb.StorageConfig
	endpoint      string
	directory     string
}

type option func(*StorageService)

func WithStorageConfig(storageConfig *storagePb.StorageConfig) option {
	return func(s *StorageService) {
		s.storageConfig = storageConfig
	}
}

func WithEndpoint(location string) option {
	return func(s *StorageService) {
		s.endpoint = location
	}
}

func WithDirectory(directory string) option {
	return func(s *StorageService) {
		s.directory = directory
	}
}

func NewStorageService(options ...option) *StorageService {
	s := &StorageService{
		storageConfig: &storagePb.StorageConfig{
			Driver: storagePb.StorageDriver_FS,
			Fs: &storagePb.FileSystemConfig{
				DataPath: "./tmp",
			},
		},
		directory: "",
		endpoint:  env.GetString("STORAGE_SERVICE_ADDR", "http://localhost:9500"),
	}

	for _, fn := range options {
		fn(s)
	}

	return s
}

func (s *StorageService) AccessToken() string {
	token, _ := EncodeStorageConfig(s.storageConfig)
	return token
}

func (s *StorageService) GetOutputPath(name string) string {
	if isValidURL(s.endpoint) {
		baseURL, err := url.Parse(s.endpoint)
		if err == nil {
			return baseURL.ResolveReference(&url.URL{Path: path.Join(s.directory, name)}).String()
		}
	}

	return path.Join(s.endpoint, s.directory, name)
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https")
}
