package main

import (
	"fmt"

	"net/url"
	"path"

	executor "github.com/mediaprodcast/commons/executor"
	"go.uber.org/zap"

	storagePb "github.com/mediaprodcast/proto/genproto/go/storage/v1"
	"github.com/mediaprodcast/storage/pkg/credentials"
)

type storageService struct {
	storageConfig  *storagePb.StorageConfig
	outputLocation string
}

func newStorageService() *storageService {
	return &storageService{
		storageConfig: &storagePb.StorageConfig{
			Driver: storagePb.StorageDriver_FS,
			Fs: &storagePb.FileSystemConfig{
				DataPath: "./tmp",
			},
		},
		outputLocation: "http://localhost:9500", // storage service
	}
}

func (s *storageService) useragent() string {
	ua, _ := credentials.Encode(s.storageConfig)
	return ua
}

func (s *storageService) getOutputPath(name string) string {
	if isValidURL(s.outputLocation) {
		baseURL, err := url.Parse(s.outputLocation)
		if err == nil {
			return baseURL.ResolveReference(&url.URL{Path: name}).String()
		}
	}

	return path.Join(s.outputLocation, name)
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https")
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	storage := newStorageService()

	args := executor.NewArgsBuilder().
		With("-i", "https://storage.googleapis.com/shaka-streamer-assets/sample-inputs/Sintel.2010.4k.mkv").
		With("-t", "2"). // Stop processing after 2 seconds
		With("-c:v", "libx264").
		With("-f", "mp4").
		With("-loglevel", "error").
		With("-movflags", "frag_keyframe+empty_moov").
		With("-method", "PUT").
		With("-headers", fmt.Sprintf("Authorization: Bearer %s", storage.useragent())).
		With(storage.getOutputPath("thumbnails/output.mp4")).
		Build()

	exec := executor.NewExecutorBuilder().
		WithBinaryPath("ffmpeg").
		WithArgs(args).
		WithLogger(logger).
		Build()

	cmd, err := exec.PreviewCommand()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(cmd)

	exec.Run()

	fmt.Println("Done Processing!")
}
