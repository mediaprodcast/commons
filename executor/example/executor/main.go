package main

import (
	"context"
	"fmt"
	"io"

	executor "github.com/mediaprodcast/commons/executor"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

// Uploads video stream to MinIO
func uploadToMinIO(bucket, key string, data io.Reader) error {
	// Initialize MinIO client
	minioClient, err := minio.New("play.min.io", &minio.Options{
		Creds:  credentials.NewStaticV4("Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", ""),
		Secure: true, // Set to true if using HTTPS
	})
	if err != nil {
		return fmt.Errorf("failed to create MinIO client: %v", err)
	}

	// Upload the video stream to the MinIO bucket
	_, err = minioClient.PutObject(context.Background(), bucket, key, data, -1, minio.PutObjectOptions{
		ContentType: "video/mp4",
	})

	if err != nil {
		return fmt.Errorf("failed to upload to MinIO: %v", err)
	}

	fmt.Println("Upload successful!")
	return nil
}

func stdoutHandler(stdout io.Reader) error {
	fmt.Println("Staring upload")
	err := uploadToMinIO("my-test-bucket", "output.mp4", stdout)
	if err != nil {
		fmt.Printf("failed to upload video to MinIO: %s\n", err.Error())
	}

	return nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	args := executor.NewArgsBuilder().
		With("-i", "https://storage.googleapis.com/shaka-streamer-assets/sample-inputs/Sintel.2010.4k.mkv").
		With("-t", "2"). // Stop processing after 2 seconds
		With("-threads", "4").
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

	cmd, err := exec.PreviewCommand()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(cmd)

	exec.Run()

	fmt.Println("Done Processing!")
}
