package main

import (
	"context"
	"fmt"
	"io"
	"log"

	storageApi "github.com/mediaprodcast/commons/api/storage"
	executor "github.com/mediaprodcast/commons/executor"
	"go.uber.org/zap"
)

var storage = storageApi.NewStorageService()

func stdoutHandler(reader io.Reader) error {
	fmt.Println("Starting upload")

	err := storage.Upload(context.Background(), "test/test-direct.mp4", reader)
	if err != nil {
		log.Println("Error uploading file:", err)
		return err
	}

	return nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	args := executor.NewArgsBuilder().
		With("-loglevel", "error").
		With("-i", "https://storage.googleapis.com/shaka-streamer-assets/sample-inputs/Sintel.2010.4k.mkv").
		With("-t", "2"). // Stop processing after 2 seconds
		With("-c:v", "libx264").
		With("-f", "mp4").
		With("-movflags", "frag_keyframe+empty_moov").
		With("pipe:1").
		Build()

	exec := executor.NewExecutorBuilder().
		WithBinaryPath("ffmpeg").
		WithArgs(args).
		WithLogger(logger).
		WithStdoutHandler(stdoutHandler).
		Build()

	if err := exec.Run(); err == nil {
		fmt.Println("Upload Processing!")
	}
}
