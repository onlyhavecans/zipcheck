package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFilesRecursive(directory string, extension string) ([]string, error) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory %s does not exist", directory)
	}

	var allZipFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == extension {
			allZipFiles = append(allZipFiles, path)
		}
		return nil
	})
	return allZipFiles, err
}
