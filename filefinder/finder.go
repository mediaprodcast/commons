package filefinder

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	storageClient "github.com/mediaprodcast/commons/api/storage"
)

type fileFinder struct {
	storage   *storageClient.StorageService
	directory string
	sorted    bool
}

func WithDirectory(directory string) func(*fileFinder) {
	return func(f *fileFinder) {
		f.directory = directory
	}
}

func WithStorageService(storage *storageClient.StorageService) func(*fileFinder) {
	return func(f *fileFinder) {
		f.storage = storage
	}
}

func WithSorted(sorted bool) func(*fileFinder) {
	return func(f *fileFinder) {
		f.sorted = sorted
	}
}

// FileFinder creates a new instance of the fileFinder struct.
// It accepts optional parameters to set the directory, storage service, and sorting flag.
// Returns a pointer to the fileFinder instance.
func FileFinder(opts ...func(*fileFinder)) *fileFinder {
	fileF := &fileFinder{}
	for _, opt := range opts {
		opt(fileF)
	}
	return fileF
}

// List retrieves a list of file paths based on the provided patterns.
// It first checks if any patterns are provided and uses the first pattern if available.
// Depending on the storage type (remote or local), it fetches the file paths accordingly.
// If the 'sorted' flag is set, it sorts the file paths numerically based on the extracted number from the file names.
// Returns a slice of file paths and an error if any occurred during the process.
//
// Parameters:
//
//	patterns - Optional variadic parameter to specify patterns for file matching.
//
// Returns:
//
//	[]string - A slice of file paths matching the provided pattern.
//	error - An error if any occurred during the file retrieval process.
func (f *fileFinder) List(patterns ...string) ([]string, error) {
	var files []string
	var err error
	var pattern string
	if len(patterns) > 0 {
		pattern = patterns[0]
	}
	if f.storage != nil {
		files, err = f.fromRemoteStorage(pattern)
	} else {
		files, err = f.fromLocalStorage(pattern)
	}
	if f.sorted {
		sort.Slice(files, func(i, j int) bool {
			numI, _ := extractNumber(filepath.Base(files[i]))
			numJ, _ := extractNumber(filepath.Base(files[j]))
			return numI < numJ
		})
	}
	return files, err
}

// fromLocalStorage retrieves a list of file paths from the local storage based on the provided pattern.
// It uses the 'filepath.Glob' function to match the pattern and retrieve the file paths.
// Returns a slice of file paths and an error if any occurred during the process.
//
// Parameters:
//
//	pattern - The pattern to match the filenames.
//
// Returns:
//
//	[]string - A slice of file paths matching the provided pattern.
//	error - An error if any occurred during the file retrieval process.
func (f *fileFinder) fromLocalStorage(pattern string) ([]string, error) {
	return filepath.Glob(filepath.Join(f.directory, pattern))
}

// fromRemoteStorage retrieves a list of file paths from the remote storage based on the provided pattern.
// It lists the files from the remote storage and matches the pattern to filter the files.
// Returns a slice of file paths and an error if any occurred during the process.
//
// Parameters:
//
//	pattern - The pattern to match the filenames.
//
// Returns:
//
//	[]string - A slice of file paths matching the provided pattern.
//	error - An error if any occurred during the file retrieval process.
func (f *fileFinder) fromRemoteStorage(pattern string) ([]string, error) {
	files := make([]string, 0)
	remoteFiles, err := f.storage.List(context.Background(), f.directory)
	if err != nil {
		return files, fmt.Errorf("failed to list files: %v", err)
	}
	for _, file := range remoteFiles {
		if match, err := filepath.Match(pattern, file.Name); err != nil || !match {
			continue
		}
		fullPath := path.Join(f.directory, file.Name)
		files = append(files, f.storage.GetAbsolutePath(fullPath))
	}
	return files, nil
}

// extractNumber extracts the number from the provided filename.
// It uses a regular expression to match the number in the filename.
// Returns the extracted number and an error if any occurred during the process.
//
// Parameters:
//
//	filename - The filename from which to extract the number.
//
// Returns:
//
//	int - The extracted number from the filename.
//	error - An error if any occurred during the extraction process.
func extractNumber(filename string) (int, error) {
	// re is a compiled regular expression that matches strings containing one or more digits
	// followed by a period and any other characters. The digits are captured in a group.
	//
	// Example matches:
	// "123.txt" -> captures "123"
	// "456.jpg" -> captures "456"
	// "789" -> no match (no period)
	// ".123" -> no match (no digits before period)
	re := regexp.MustCompile(`(\d+)\..*`)
	match := re.FindStringSubmatch(filename)
	if match == nil {
		return 0, fmt.Errorf("invalid filename format: %s", filename)
	}
	idStr := match[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert ID to integer: %w", err)
	}
	return id, nil
}
