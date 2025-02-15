package main

import (
	"fmt"

	storageApi "github.com/mediaprodcast/commons/api/storage"
	executor "github.com/mediaprodcast/commons/executor"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	storage := storageApi.NewStorageService()
	args := executor.NewArgsBuilder().
		With("-i", "https://storage.googleapis.com/shaka-streamer-assets/sample-inputs/Sintel.2010.4k.mkv").
		With("-t", "2").
		With("-c:v", "libx264").
		With("-f", "mp4").
		With("-loglevel", "error").
		With("-movflags", "frag_keyframe+empty_moov").
		With("-method", "PUT").
		With("-headers", fmt.Sprintf("Authorization: Bearer %s", storage.AccessToken())).
		With(storage.GetOutputPath("thumbnails/output.mp4")).
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
