package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/mediaprodcast/proto/genproto/go/storage/v1"
	"go.uber.org/zap"
)

func (s *StorageService) Stats(ctx context.Context, path string) error {
	panic("implement me")
}

func (s *StorageService) List(ctx context.Context, name string) ([]*pb.Stat, error) {
	s.logger.Debug("Starting listing", zap.String("path", name))
	files := []*pb.Stat{}

	req, err := http.NewRequestWithContext(ctx, "GET", s.GetAbsolutePath(name), nil)
	if err != nil {
		s.logger.Error("Failed to create request", zap.String("path", name), zap.Error(err)) // Log failure to create request
		return files, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		s.logger.Error("Failed to send request", zap.String("path", name), zap.Error(err)) // Log failure to send request
		return files, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent { // 204 No Content is also a valid list response
		s.logger.Error("list failed", zap.String("path", name), zap.Int("statusCode", resp.StatusCode)) // Log failed list response
		return files, fmt.Errorf("list failed with status: %s", resp.Status)
	}

	s.logger.Debug("Listed successful", zap.String("path", name))

	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		s.logger.Error("Failed to decode response", zap.String("path", name), zap.Error(err)) // Log failure to decode response
		return files, fmt.Errorf("failed to decode response: %v", err)
	}

	return files, nil
}
