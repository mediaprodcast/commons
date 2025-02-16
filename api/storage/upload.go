package storage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

// uploadTask represents a single file upload task.
type uploadTask struct {
	name   string
	reader io.Reader
}

// Upload uploads a single file to the storage service.
func (s *StorageService) Upload(ctx context.Context, name string, reader io.Reader, contentType ...string) error {
	s.logger.Info("Starting upload", zap.String("file", name)) // Log starting upload

	req, err := http.NewRequestWithContext(ctx, "PUT", s.GetAbsolutePath(name), reader)
	if err != nil {
		s.logger.Error("Failed to create request", zap.String("file", name), zap.Error(err)) // Log failure to create request
		return fmt.Errorf("failed to create request: %v", err)
	}

	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType[0])
	}

	resp, err := s.client.Do(req)
	if err != nil {
		s.logger.Error("Failed to send request", zap.String("file", name), zap.Error(err)) // Log failure to send request
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Error("Upload failed", zap.String("file", name), zap.Int("statusCode", resp.StatusCode)) // Log failed upload
		return fmt.Errorf("upload failed with status: %s", resp.Status)
	}

	s.logger.Info("Upload successful", zap.String("file", name)) // Log successful upload
	return nil
}

// UploadBulk uploads multiple files concurrently using workers.
func (s *StorageService) UploadBulk(ctx context.Context, files map[string]io.Reader) error {
	var wg sync.WaitGroup
	fileChan := make(chan uploadTask, len(files))
	errChan := make(chan error, len(files))

	// Start worker goroutines.
	for i := 0; i < s.workers; i++ {
		wg.Add(1)
		go s.uploadWorker(ctx, &wg, fileChan, errChan)
	}

	// Send files to the workers.
	for name, reader := range files {
		s.logger.Info("Queueing file for upload", zap.String("file", name)) // Log queueing file
		fileChan <- uploadTask{name: name, reader: reader}
	}
	close(fileChan)

	// Wait for all workers to finish.
	wg.Wait()
	close(errChan)

	// Collect errors from the error channel.
	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		s.logger.Error("Failed to upload multiple files", zap.Int("failedCount", len(errors)), zap.Error(errors[0])) // Log multiple upload failure
		return fmt.Errorf("failed to upload %d files: %v", len(errors), errors)
	}

	s.logger.Info("All files uploaded successfully") // Log all files uploaded successfully
	return nil
}

// uploadWorker processes upload tasks from the channel.
func (s *StorageService) uploadWorker(ctx context.Context, wg *sync.WaitGroup, fileChan <-chan uploadTask, errChan chan<- error) {
	defer wg.Done()

	for task := range fileChan {
		s.logger.Debug("Starting worker upload", zap.String("file", task.name)) // Log starting worker upload
		if err := s.Upload(ctx, task.name, task.reader); err != nil {
			s.logger.Error("Failed to upload file in worker", zap.String("file", task.name), zap.Error(err)) // Log failure in worker
			errChan <- fmt.Errorf("failed to upload file %s: %v", task.name, err)
		} else {
			s.logger.Debug("Worker finished upload", zap.String("file", task.name)) // Log finished upload in worker
		}
	}
}
