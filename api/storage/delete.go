package storage

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

// Delete deletes a single file from the storage service.
func (s *StorageService) Delete(ctx context.Context, name string) error {
	s.logger.Info("Starting delete", zap.String("file", name)) // Log starting delete

	req, err := http.NewRequestWithContext(ctx, "DELETE", s.GetAbsolutePath(name), nil)
	if err != nil {
		s.logger.Error("Failed to create request", zap.String("file", name), zap.Error(err)) // Log failure to create request
		return fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		s.logger.Error("Failed to send request", zap.String("file", name), zap.Error(err)) // Log failure to send request
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent { // 204 No Content is also a valid delete response
		s.logger.Error("Delete failed", zap.String("file", name), zap.Int("statusCode", resp.StatusCode)) // Log failed delete response
		return fmt.Errorf("delete failed with status: %s", resp.Status)
	}

	s.logger.Info("Delete successful", zap.String("file", name)) // Log successful delete
	return nil
}

// DeleteBulk deletes multiple files concurrently using workers.
func (s *StorageService) DeleteBulk(ctx context.Context, files []string) error {
	var wg sync.WaitGroup
	nameChan := make(chan string, len(files))
	errChan := make(chan error, len(files))

	// Start worker goroutines.
	for i := 0; i < s.workers; i++ {
		wg.Add(1)
		go s.deleteWorker(ctx, &wg, nameChan, errChan)
	}

	// Send file names to the workers.
	for _, name := range files {
		s.logger.Info("Queueing file for deletion", zap.String("file", name)) // Log queuing file for deletion
		nameChan <- name
	}
	close(nameChan)

	// Wait for all workers to finish.
	wg.Wait()
	close(errChan)

	// Collect errors from the error channel.
	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		s.logger.Error("Failed to delete multiple files", zap.Int("failedCount", len(errors)), zap.Error(errors[0])) // Log multiple delete failures
		return fmt.Errorf("failed to delete %d files: %v", len(errors), errors)
	}

	s.logger.Info("All files deleted successfully") // Log all files deleted successfully
	return nil
}

// deleteWorker processes delete tasks from the channel.
func (s *StorageService) deleteWorker(ctx context.Context, wg *sync.WaitGroup, nameChan <-chan string, errChan chan<- error) {
	defer wg.Done()

	for name := range nameChan {
		s.logger.Debug("Starting worker delete", zap.String("file", name)) // Log starting worker delete
		if err := s.Delete(ctx, name); err != nil {
			s.logger.Error("Failed to delete file in worker", zap.String("file", name), zap.Error(err)) // Log failure in worker
			errChan <- fmt.Errorf("failed to delete file %s: %v", name, err)
		} else {
			s.logger.Debug("Worker finished delete", zap.String("file", name)) // Log worker finish
		}
	}
}
