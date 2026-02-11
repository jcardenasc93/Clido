package storage

import "fmt"

type LoadFileError struct {
	filePath string
	details  string
}

func (e *LoadFileError) Error() string {
	return fmt.Sprintf("Couldn't load content from source: %s\n%s\n", e.filePath, e.details)
}
